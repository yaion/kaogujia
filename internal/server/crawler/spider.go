package crawler

import (
	"github.com/gocolly/colly/v2"
	"kaogujia/pkg/config"
	"kaogujia/pkg/utils"
	"log"
	"time"
)

type Spider struct {
	collector *colly.Collector
}

func NewSpider(website config.WebsiteConfig) *Spider {
	c := colly.NewCollector(
		colly.AllowedDomains(website.AllowedDomains...),
		colly.Async(true),
	)

	// 使用网站特定的请求间隔配置
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: website.Parallelism,
		RandomDelay: time.Duration(website.RandomDelay) * time.Second,
	})

	return &Spider{collector: c}
}

func (s *Spider) SetupHandlers() {
	// 错误处理
	s.collector.OnError(func(r *colly.Response, err error) {
		log.Printf("请求失败: %s | 错误: %v", r.Request.URL, err)
	})

	// 响应处理
	s.collector.OnResponse(func(r *colly.Response) {
		log.Printf("收到响应: %d | %s", r.StatusCode, r.Request.URL)

		// 解密示例（如果需要）
		decrypted, err := utils.Decrypt(r.Request.URL.String(), string(r.Body))
		if err == nil {
			log.Println("解密后的数据:", decrypted)
		}
	})

	// 处理特定选择器
	s.collector.OnHTML(".product-item", func(e *colly.HTMLElement) {
		// 提取数据
		name := e.ChildText(".name")
		price := e.ChildText(".price")

		log.Printf("产品: %s, 价格: %s", name, price)

		// TODO: 保存到数据库
	})
}

func (s *Spider) Start(urls []string) {
	for _, url := range urls {
		s.collector.Visit(url)
	}
	s.collector.Wait()
}
