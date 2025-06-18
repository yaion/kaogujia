package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func decrypt(urlStr, text string) (string, error) {
	// 1. 参数检查
	if urlStr == "" || text == "" {
		return "", fmt.Errorf("URL and text must not be empty")
	}

	// 2. URL编码
	str := getStr(urlStr)

	// 5. 提取密钥和IV
	orgKey := str[:16]
	orgIV := str[12:28]

	// 6. 解码Base64密文
	ciphertext, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", fmt.Errorf("base64 decode error: %v", err)
	}

	// 7. 创建AES解密器
	block, err := aes.NewCipher([]byte(orgKey))
	if err != nil {
		return "", fmt.Errorf("cipher creation error: %v", err)
	}

	// 8. 检查IV长度
	if len(orgIV) != block.BlockSize() {
		return "", fmt.Errorf("IV length must equal block size")
	}

	// 9. 创建CBC模式解密器
	mode := cipher.NewCBCDecrypter(block, []byte(orgIV))

	// 10. 解密数据
	decrypted := make([]byte, len(ciphertext))
	mode.CryptBlocks(decrypted, ciphertext)

	// 11. 移除PKCS7填充
	unpadded, err := unpadPKCS7(decrypted)
	if err != nil {
		return "", fmt.Errorf("padding removal error: %v", err)
	}

	// 12. 返回解密后的字符串
	return string(unpadded), nil
}

// 移除PKCS7填充
func unpadPKCS7(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("input data is empty")
	}

	padding := int(data[len(data)-1])
	if padding < 1 || padding > aes.BlockSize {
		return nil, fmt.Errorf("invalid padding size")
	}

	if len(data) < padding {
		return nil, fmt.Errorf("data length is less than padding size")
	}

	// 检查填充字节是否有效
	for i := len(data) - padding; i < len(data); i++ {
		if int(data[i]) != padding {
			return nil, fmt.Errorf("invalid padding byte")
		}
	}

	return data[:len(data)-padding], nil
}

func getStr(url string) string {
	// 1. 执行encodeURI
	encoded := encodeURI(url)

	// 2. 进行Base64编码
	base64Str := base64.StdEncoding.EncodeToString([]byte(encoded))

	// 3. 重复3次并返回
	return strings.Repeat(base64Str, 3)
}

func encodeURI(s string) string {
	safe := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_.!~*'();,/?:@&=+$#"
	var result strings.Builder
	for _, r := range s {
		if strings.ContainsRune(safe, r) {
			result.WriteRune(r)
		} else {
			// UTF-8编码非安全字符
			for _, b := range []byte(string(r)) {
				result.WriteString(fmt.Sprintf("%%%02X", b))
			}
		}
	}
	return result.String()
}

func main() {
	mapURL := make(map[string]string{}, 0)
	// 示例用法
	// 达人 /api/author/search
	// - 对于视频博主，"video_ratio": 1, "live_ratio": 0
   //- 对于直播博主，"live_ratio": 1, "video_ratio": 0
	//urlApi := "https://service.kaogujia.com/api/author/search?limit=50&page=1&sort_field=gmv&sort=0"
	//urlApi := "https://service.kaogujia.com/api/author/search?limit=50&page=1&sort_field=live_gmv&sort=0"
	urlApi := "https://service.kaogujia.com/api/author/search?limit=50&page=10&sort_field=video_gmv&sort=0"
	// 把json输出未一个文件
	//filename := "author.json"
	//filename := "author_live_gmv.json"
	filename := "author_video_gmv.json"

	//params := `{"keyword":"","author_type":0}` // 关键词和作者类型
	//params := `{"keyword":"","author_type":1}`
	params := `{"keyword":"","author_type":2}`

	//商品

	// 直播

	// 视频

	//小店

	//品牌

	parsedURL, err := url.Parse(urlApi)
	if err != nil {
		log.Output(1, fmt.Sprintf("URL parse error: %v", err))
		return
	}
	url := parsedURL.Path

	encryptedText, err := GetApi(urlApi, params)
	if err != nil {
		log.Output(1, fmt.Sprintf("API call error: %v", err))
		return
	}

	result, err := decrypt(url, encryptedText)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}

	res := make(map[string]interface{})
	json.Unmarshal([]byte(result), &res)

	jsStr, _ := json.Marshal(res)

	err = writeToFile(filename, string(jsStr))
	if err != nil {
		panic(err)
	}

}

func writeToFile(filename, content string) error {
	// 创建或打开文件（如果文件不存在则创建）
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 使用 bufio 提高写入效率
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		return err
	}

	// 将缓冲区中的内容刷新到文件中
	err = writer.Flush()
	if err != nil {
		return err
	}

	//fmt.Println("写入成功！")
	return nil
}
