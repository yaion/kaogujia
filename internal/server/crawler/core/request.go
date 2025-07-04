package core

import (
	"bytes"
	"context"
	"github.com/gocolly/colly"
	"io"
	"log"
	"net/http"
	"time"
)

type Task struct {
	URL     string
	Method  string
	Headers map[string]string
	Body    []byte
	Handler func(*colly.Response, *Account) error
	Meta    map[string]interface{}
}

type TaskSpec struct {
	URL     string            `yaml:"url"`
	Method  string            `yaml:"method"`
	Headers map[string]string `yaml:"headers"`
	Body    []byte            `yaml:"body"`
	Handler string            `yaml:"handler"`
	Meta    map[string]interface{}
}

type ResponseHandler func(*colly.Response, *Account) error

func ExecuteRequest(task *Task, account *Account, dispatcher *TaskDispatcher) error {
	// 应用速率限制
	if account.RateLimit != nil {
		start := time.Now()
		account.RateLimit.Wait()
		log.Printf("账号 %s 请求等待: %v", account.ID, time.Since(start))
	}

	c := colly.NewCollector(
		colly.Async(true),
		colly.IgnoreRobotsTxt(),
	)

	// 设置超时
	c.SetRequestTimeout(30 * time.Second)

	// 设置代理
	if account.Proxy != "" {
		if err := c.SetProxy(account.Proxy); err != nil {
			log.Printf("设置代理失败: %v", err)
			return err
		}
	}

	/*// 设置cookies
	for _, cookie := range account.Cookies {
		if ck, ok := cookie.(*http.Cookie); ok {
			c.SetCookie(ck.String())
		}
	}*/

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

	// 注册响应处理
	c.OnResponse(func(r *colly.Response) {
		if err := task.Handler(r, account, dispatcher); err != nil {
			log.Printf("处理器错误: %v", err)
		}
	})

	// 处理请求错误
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("请求失败: %s, 错误: %v", task.URL, err)
	})

	// 在请求前设置元数据到上下文
	c.OnRequest(func(r *colly.Request) {
		ctx := context.WithValue(r.Context(), "meta", task.Meta)
		r.Request = r.Request.WithContext(ctx)
	})

	return c.Request(request.Method, request.URL.String(), request.Body, nil, request.Header)
}
