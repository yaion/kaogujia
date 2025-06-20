package utils

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClient(proxyURL string) (*HttpClient, error) {
	// 设置代理ip
	transport := new(http.Transport)
	if proxyURL != "" {
		proxy, err := url.Parse(proxyURL)
		if err != nil {
			log.Printf("Error parsing proxy URL: %v", err)
			return nil, err
		}
		transport.Proxy = http.ProxyURL(proxy)
		// 设置代理连接池大小
		transport.MaxIdleConns = 10
		transport.MaxConnsPerHost = 10
	}

	return &HttpClient{
		client: &http.Client{
			Transport: transport,
		},
	}, nil
}

// SendRequest 公共HTTP请求方法
// method: 请求方法 (GET, POST, PUT, DELETE等)
// url: 请求URL
// body: 请求体内容
// headers: 自定义请求头 (可选)
func (c *HttpClient) SendRequest(method, url string, body []byte, headers map[string]string) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// Get 使用公共方法实现的GET请求
func (c *HttpClient) Get(url string, headers map[string]string) ([]byte, error) {
	return c.SendRequest(http.MethodGet, url, nil, headers)
}

// Post 使用公共方法实现的POST请求
func (c *HttpClient) Post(url string, body []byte, headers map[string]string) ([]byte, error) {
	return c.SendRequest(http.MethodPost, url, body, headers)
}
