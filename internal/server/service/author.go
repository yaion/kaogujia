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

func GetAuthorList() (map[string]interface{}, error) {
	client := middleware.GetMongo()
	db := client.Database("kaogujia")
	authorDAO := mongodb.NewAuthorDAO(db)
	result, err := authorDAO.ListAll(context.Background(), bson.M{}, 1, 10)
	if err != nil {
		log.Printf("Get author list error: %v", err)
		return nil, err
	}
	return result, err
}
