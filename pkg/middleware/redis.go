package middleware

import (
	"context"
	"fmt"
	"goWebBasicProject/pkg/config"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() error {
	conf := config.Get().Redis
	addr := fmt.Sprintf("%s:%v", conf.Addr, conf.Port)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Password,
		DB:       conf.DB,
		PoolSize: 20,
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := RedisClient.Ping(ctx).Result()
	return err
}
