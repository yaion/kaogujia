package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"kaogujia/internal/server/dto/request"
	"kaogujia/internal/server/dto/response"
	"kaogujia/internal/server/service"
)

// GetAuthorList 搜索达人信息 达人列表接口
func GetAuthorList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req request.AuthorSearchRequest
	err = c.BindAndValidate(&req)
	resp := new(response.AuthorSearchResponse)
	if err != nil {
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	params := make(map[string]interface{})

	result, err := service.GetAuthorList(ctx, params, req.Page, req.Limit)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	resp.Code = "200"
	resp.Data = result
	c.JSON(consts.StatusOK, resp)
}

// GetAuthorInfo 获取达人信息
func GetAuthorInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req request.AuthorInfoRequest
	err = c.BindAndValidate(&req)
	resp := new(response.AuthorSearchResponse)
	if err != nil {
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	result, err := service.GetAuthorInfo(ctx, req.Uid)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	resp.Code = "200"
	resp.Data = result
	c.JSON(consts.StatusOK, resp)
}
