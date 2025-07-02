package middleware

import (
	"log"
)

// InitAll 初始化所有中间件连接
func InitAll() error {
	// 按依赖顺序初始化
	/*if err := InitMySQL(); err != nil {
		log.Fatalf("MySQL init failed: %v", err)
		return err
	}

	if err := InitRedis(); err != nil {
		log.Fatalf("Redis init failed: %v", err)
		return err
	}*/

	if err := InitMongo(); err != nil {
		log.Fatalf("MongoDB init failed: %v", err)
		return err
	}

	log.Println("All middleware initialized successfully")
	return nil
}

// CloseAll 关闭所有连接
func CloseAll() {
	/*if sqlDB, _ := DB.DB(); sqlDB != nil {
		sqlDB.Close()
	}

	if RedisClient != nil {
		RedisClient.Close()
	}*/
}
