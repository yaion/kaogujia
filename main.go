package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"kaogujia/pkg/utils"
	"log"
	"net/url"
	"os"
)

func main() {

	// 示例用法
	//url := "/api/rank/live/sku/author"
	urlApi := "https://service.kaogujia.com/api/author/search?limit=50&page=1&sort_field=gmv&sort=0"

	parsedURL, err := url.Parse(urlApi)
	if err != nil {
		fmt.Println(err)
		return
	}
	url := parsedURL.Path

	encryptedText, err := GetApi(urlApi)
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := utils.Decrypt(url, encryptedText)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}

	fmt.Println("Decrypted Result:")
	//fmt.Println(result)

	res := make(map[string]interface{})
	json.Unmarshal([]byte(result), &res)
	fmt.Println("-------------------")
	//fmt.Println(res)

	jsStr, _ := json.Marshal(res)
	fmt.Println("---------json----------")
	fmt.Println(string(jsStr))
	// 把json输出未一个文件
	filename := "output1.txt"

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

	fmt.Println("写入成功！")
	return nil
}
