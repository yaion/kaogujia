package main

import (
	"bufio"
	"fmt"
	serve "kaogujia/cmd/server"
	"os"
)

func main() {
	serve.Run()
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
