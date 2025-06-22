package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"kaogujia/internal/server/dto/request"
	"kaogujia/internal/server/dto/response"
)

// GetAuthor 搜索达人信息 达人列表接口
func GetAuthor(ctx context.Context, c *app.RequestContext) {
	var err error
	var req request.AuthorSearchRequest
	err = c.BindAndValidate(&req)
	resp := new(response.AuthorSearchResponse)
	if err != nil {
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	/*if err = mysql.CreateUser([]*model.User{
		{
			Name:      req.Name,
			Gender:    int64(req.Gender),
			Age:       req.Age,
			Introduce: req.Introduce,
		},
	}); err != nil {
		c.JSON(consts.StatusInternalServerError, &user_gorm.CreateUserResponse{Code: user_gorm.Code_DBErr, Msg: err.Error()})
		return
	}*/

	//resp := new(response.AuthorSearchResponse)
	resp.Code = "200"
	resp.Data = req
	c.JSON(consts.StatusOK, resp)
}
