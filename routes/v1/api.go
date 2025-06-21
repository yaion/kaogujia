package v1

import "github.com/cloudwego/hertz/pkg/app/server"

func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_v1 := root.Group("/v1", _v1Mw()...)
		{
			_user := _v1.Group("/user", _userMw()...)
			{
				_create := _user.Group("/create", _createMw()...)
				_create.POST("/", append(_createuserMw(), user_gorm.CreateUser)...)
			}
			{
				_delete := _user.Group("/delete", _deleteMw()...)
				_delete.POST("/:user_id", append(_deleteuserMw(), user_gorm.DeleteUser)...)
			}
			{
				_query := _user.Group("/query", _queryMw()...)
				_query.POST("/", append(_queryuserMw(), user_gorm.QueryUser)...)
			}
			{
				_update := _user.Group("/update", _updateMw()...)
				_update.POST("/:user_id", append(_updateuserMw(), user_gorm.UpdateUser)...)
			}
		}
	}
}
