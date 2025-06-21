package server

import (
	"kaogujia/internal/server/crawler"
	"kaogujia/pkg/config"
	"kaogujia/routes"

	app "github.com/cloudwego/hertz/pkg/app/server"
)

type Server struct {
	crawler *crawler.Scheduler // 添加爬虫调度器
}

func NewServer(cfg *config.AppConfig) *Server {
	serve := new(Server)
	serve.crawler = crawler.NewScheduler(cfg)

	return serve
}

func (s *Server) Run() {
	s.Crawler()
	s.web()
}

func (s *Server) Crawler() {
	s.crawler.Run()
}

func (s *Server) web() {
	h := app.Default()
	routes.RegisterRoutes(h)

	h.Spin()
}
