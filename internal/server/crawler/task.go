package crawler

import (
	"kaogujia/pkg/config"
	"time"
)

type Scheduler struct {
	interval time.Duration
}

func NewScheduler(cfg *config.AppConfig) *Scheduler {
	return &Scheduler{
		interval: time.Minute * time.Duration(cfg.Crawler.Interval),
	}
}

func (s *Scheduler) Run() {
	// 结合定时任务
	ticker := time.NewTicker(s.interval)
	for range ticker.C {
		// 执行爬虫任务
		//s.ExecuteTask()
	}
}
