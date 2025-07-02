package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"kaogujia/internal/server/dao/mongodb"
	"kaogujia/pkg/middleware"
	"log"
)

func GetProductList(ctx context.Context, params map[string]interface{}, page, limit int64) (map[string]interface{}, error) {
	client := middleware.GetMongo()
	db := client.Database("kaogujia")
	DAO := mongodb.NewProductDAO(db)
	bs := bson.M(params)
	result, err := DAO.ListAll(ctx, bs, page, limit)
	if err != nil {
		log.Printf("Get product list error: %v", err)
		return nil, err
	}
	return result, err
}
