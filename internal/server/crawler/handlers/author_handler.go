package handlers

import (
	"encoding/json"
	"github.com/gocolly/colly/v2"
	"kaogujia/internal/server/crawler/core"
	"kaogujia/internal/server/dao/mongodb"
	"kaogujia/internal/server/service"
)

type AuthorResult struct {
	IsAuthority bool              `json:"is_authority"`
	Items       []*mongodb.Author `json:"items"`
	Pagination  Pagination        `json:"pagination"`
	Sort        Sort              `json:"sort"`
}

func AuthorHandler(r *colly.Response, acc *core.Account) error {
	str, err := Handler(r)
	if err != nil {
		return err
	}
	result := new(AuthorResult)
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return err
	}
	// 记录当前放回的数据
	err = service.CreateAuthors(result.Items)
	if err != nil {
		return err
	}
	// todo 根据单曲列表数量判断是否继续爬取
	if len(result.Items) > 0 {
	}

	return nil
}
