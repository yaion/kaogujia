package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"kaogujia/internal/server/dao/mongodb"
	"kaogujia/pkg/middleware"
	"log"
)

func CreateAuthors(authors []*mongodb.Author) error {
	client := middleware.GetMongo()
	db := client.Database("kaogujia")
	fmt.Println("db:", db)
	authorDAO := mongodb.NewAuthorDAO(db)
	var docs []interface{}
	for _, author := range authors {
		docs = append(docs, author)
	}
	err := authorDAO.BatchCreate(context.Background(), docs)
	if err != nil {
		log.Printf("Create author error: %v", err)
		return err
	}
	log.Println("Create author success")
	return nil
}

func GetAuthorList(ctx context.Context, params map[string]interface{}, page, limit int64) (map[string]interface{}, error) {
	client := middleware.GetMongo()
	db := client.Database("kaogujia")
	authorDAO := mongodb.NewAuthorDAO(db)
	bs := bson.M(params)
	result, err := authorDAO.ListAll(ctx, bs, page, limit)
	if err != nil {
		log.Printf("Get author list error: %v", err)
		return nil, err
	}
	return result, err
}

func GetAuthorInfo(ctx context.Context, uid string) (*mongodb.AuthorInfo, error) {
	client := middleware.GetMongo()
	db := client.Database("kaogujia")
	authorInfoDAO := mongodb.NewAuthorInfo(db)
	info, err := authorInfoDAO.GetByID(ctx, uid)
	if err != nil {
		log.Printf("Get author info error: %v", err)
		return nil, err
	}
	return info, err
}
