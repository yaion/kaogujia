package crawler

import (
	"kaogujia/pkg/config"
	"time"
)

type Scheduler struct {
	interval time.Duration
	spider   *Spider
}

func NewScheduler(cfg *config.AppConfig) *Scheduler {
	spider := NewSpider()
	spider.SetupHandlers()

	return &Scheduler{
		interval: time.Minute * time.Duration(cfg.Crawler.Interval),
		spider:   spider,
	}
}

func (s *Scheduler) Run() {
	// 立即执行一次
	s.ExecuteTask()

	// 定时执行
	ticker := time.NewTicker(s.interval)
	for range ticker.C {
		s.ExecuteTask()
	}
}

func (s *Scheduler) ExecuteTask() {
	// 从配置获取目标URL
	//cfg := config.Get()
	urls := []string{"", ""}
	s.spider.Start(urls)
	account := AccountManager.GetAccount()

	// 设置专属代理
	proxyClient := utils.NewHttpClientWithProxyPool([]string{account.Proxy})

	// 设置请求间隔
	time.Sleep(account.RateLimit)

	// 使用带账号状态的客户端请求
	for _, url := range s.crawler.Targets {
		headers := map[string]string{
			"Cookie": formatCookies(account.Cookies),
		}
		proxyClient.Get(url, headers)
	}
}
