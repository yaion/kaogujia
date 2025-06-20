package serve

import (
	"kaogujia/internal/server"
	"kaogujia/pkg/config"
	"kaogujia/pkg/middleware"
	"log"
)

func serve() {
	// 1. 加载配置
	if err := config.Load("configs/app.yaml"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. 初始化中间件 (传递配置)
	cfg := config.Get()
	if err := middleware.InitAll(); err != nil {
		log.Fatalf("Middleware init failed: %v", err)
	}

	// 3. 启动服务器
	server := server.NewServer(cfg)
	server.Run()

	// 4 . 爬虫服务
	go server.Crawler()
}
