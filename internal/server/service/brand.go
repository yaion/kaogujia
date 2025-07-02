package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"kaogujia/internal/server/dao/mongodb"
	"kaogujia/pkg/middleware"
	"log"
)

func GetBrandList(ctx context.Context, params map[string]interface{}, page, limit int64) (map[string]interface{}, error) {
	client := middleware.GetMongo()
	db := client.Database("kaogujia")
	DAO := mongodb.NewBrandDAO(db)
	bs := bson.M(params)
	result, err := DAO.ListAll(ctx, bs, page, limit)
	if err != nil {
		log.Printf("Get brand list error: %v", err)
		return nil, err
	}
	return result, err
}
