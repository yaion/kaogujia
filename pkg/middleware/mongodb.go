package middleware

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"kaogujia/pkg/config"
	"log"
	"time"
)

var MongoClient *mongo.Client

func InitMongo() error {
	conf := config.Get().Mongo
	// 设置连接配置
	clientOptions := options.Client().ApplyURI(conf.Uri)

	// 建立连接（含超时控制）
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.TimeOut))
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Disconnect(ctx) // 确保关闭连接

	// 检查连接状态
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

func GetMongo() *mongo.Client {
	return MongoClient
}
