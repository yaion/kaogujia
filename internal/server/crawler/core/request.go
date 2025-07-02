package core

import (
	"bytes"
	"github.com/gocolly/colly"
	"io"
	"net/http"
)

type Task struct {
	URL     string
	Method  string
	Headers map[string]string
	Body    []byte
	Handler func(*colly.Response, *Account) error
}

type ResponseHandler func(*colly.Response, *Account) error

func ExecuteRequest(task *Task, account *Account) error {
	c := colly.NewCollector()

	// 设置代理
	if account.Proxy != "" {
		c.SetProxy(account.Proxy)
	}

	// 注册响应处理
	c.OnResponse(func(r *colly.Response) {
		if err := task.Handler(r, account); err != nil {
			// 错误处理逻辑
		}
	})

	// 创建请求
	var body io.Reader
	if task.Body != nil {
		body = bytes.NewReader(task.Body)
	}

	request, err := http.NewRequest(task.Method, task.URL, body)
	if err != nil {
		return err
	}

	// 设置请求头
	for k, v := range task.Headers {
		request.Header.Set(k, v)
	}

	return c.Request(request.Method, request.URL.String(), request.Body, nil, request.Header)
}
