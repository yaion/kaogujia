package server

import (
	"kaogujia/internal/server/crawler"
	"kaogujia/pkg/config"
)

type Server struct {
	crawler *crawler.Scheduler // 添加爬虫调度器
}

func NewServer(cfg *config.AppConfig) *Server {
	serve := new(Server)
	//todo 使用那个web框架
	serve.crawler = crawler.NewScheduler(cfg)

	return serve
}

func (s *Server) Run() {

}

func (s *Server) Crawler() {
	s.crawler.Run()
}
