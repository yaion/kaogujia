package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Result struct {
	Code    int
	Message string
	Success bool
	Data    string
}

func GetApi(url, param string) (string, error) {

	//url := "https://service.kaogujia.com/api/author/search?limit=50&page=1&sort_field=gmv&sort=0"
	method := "POST"

	payload := strings.NewReader(param) //`{"keyword":"","author_type":0}`

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Output(1, fmt.Sprintf("HTTP request error: %v", err))
		return "", err
	}
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "zh-HK,zh-CN;q=0.9,zh;q=0.8,zh-TW;q=0.7")
	req.Header.Add("authorization", "Bearer eyJhbGciOiJIUzUxMiJ9.eyJhdWQiOiIxMDAwIiwiaXNzIjoia2FvZ3VqaWEuY29tIiwianRpIjoiNjI3MmYyN2EyZDU5NDc0YThhYzk1NTQyNzgyYjM4OWIiLCJzaWQiOjgyMjY4NzcsImlhdCI6MTc0OTkwMDQ0MywiZXhwIjoxNzUwNTA1MjQzLCJid2UiOjAsInR5cCI6MSwicF9id2UiOjB9.sTul3qBenukj-HiOsTS_CnzHM0TV91cLA_U6dm6U5Z5ZFYgu6ZeTM3_Ai4AYdmvDN7q_SMoFjoQvv_LNo2VdzQ")
	req.Header.Add("origin", "https://www.kaogujia.com")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("referer", "https://www.kaogujia.com/")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"137\", \"Chromium\";v=\"137\", \"Not/A)Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36")
	req.Header.Add("version_code", "3.1")
	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Output(1, fmt.Sprintf("HTTP request error: %v", err))
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Output(1, fmt.Sprintf("HTTP request error: %v", err))
		return "", err
	}

	result := new(Result)
	err = json.Unmarshal(body, result)
	if err != nil {
		log.Output(1, fmt.Sprintf("JSON unmarshal error: %v", err))
		return "", err
	}
	return result.Data, err
}
