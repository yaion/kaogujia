package crawler

import (
	"kaogujia/pkg/config"
	"time"
)

type WebsiteScheduler struct {
	website  string
	interval time.Duration
	spider   *Spider
	targets  []string
}

type MultiScheduler struct {
	schedulers []*WebsiteScheduler
}

func NewMultiScheduler(cfg *config.AppConfig) *MultiScheduler {
	schedulers := []*WebsiteScheduler{}

	for _, website := range cfg.Websites {
		spider := NewSpider(website) // 传入网站配置
		spider.SetupHandlers()

		schedulers = append(schedulers, &WebsiteScheduler{
			website:  website.Name,
			interval: time.Minute * time.Duration(website.Interval),
			spider:   spider,
			targets:  website.Targets,
		})
	}

	return &MultiScheduler{schedulers: schedulers}
}

func (m *MultiScheduler) Run() {
	for _, s := range m.schedulers {
		go s.run()
	}
}

func (s *WebsiteScheduler) run() {
	// 立即执行一次
	//s.ExecuteTask()

	// 定时执行
	ticker := time.NewTicker(s.interval)
	for range ticker.C {
		//s.ExecuteTask()
	}
}
