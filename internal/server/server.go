package server

import (
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/gzip"
	"kaogujia/pkg/config"
	"kaogujia/routes"

	app "github.com/cloudwego/hertz/pkg/app/server"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/logger/accesslog"
)

type Server struct {
	//crawler *crawler.Scheduler // 添加爬虫调度器
}

func NewServer(cfg *config.AppConfig) *Server {
	serve := new(Server)
	//serve.crawler = crawler.NewScheduler(cfg)

	return serve
}

func (s *Server) Run() {
	s.Crawler()
	s.web()
}

func (s *Server) Crawler() {
	//s.crawler.Run()
}

func (s *Server) web() {
	//h := app.Default()
	h := app.New()
	h.Use(recovery.Recovery())
	routes.RegisterRoutes(h)

	h.Spin()
}

func registerMiddleware(h *app.Hertz) {
	cfg := config.Get()
	// log
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(config.LogLevel())
	hlog.SetOutput(
		&lumberjack.Logger{
			Filename:   cfg.Log.LogFileName,
			MaxSize:    cfg.Log.LogMaxSize,
			MaxBackups: cfg.Log.LogMaxBackups,
			MaxAge:     cfg.Log.LogMaxAge,
		},
	)

	// gzip
	if cfg.Server.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if cfg.Server.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())

}
