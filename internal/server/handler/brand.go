package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"kaogujia/internal/server/dto/request"
	"kaogujia/internal/server/service"
)

func GetBrandList(ctx context.Context, c *app.RequestContext) {
	var req request.LiveSearchRequest
	if err := c.BindAndValidate(&req); err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	params := make(map[string]interface{})
	// todo 添加搜索选项

	result, err := service.GetBrandList(ctx, params, req.Page, req.Limit)
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(consts.StatusOK, result)
}
