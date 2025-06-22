package v1

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"kaogujia/internal/server/handler"
)

func Register(r *server.Hertz) {

	root := r.Group("/api", rootMw()...)
	{
		// 达人
		author := root.Group("/author")
		{
			author.POST("/search", handler.GetAuthor)

		}

		// 商品 api/sku/search
		sku := root.Group("/sku")
		{
			sku.POST("/search", handler.GetAuthor)

		}

		// 直播列表 api/live/search

		live := root.Group("/live")
		{
			live.POST("/search", handler.GetAuthor)

		}

		// 视频 /api/video/search

		video := root.Group("/video")
		{
			video.POST("/search", handler.GetAuthor)

		}

		// 小店列表 api/shop/search

		shop := root.Group("/shop")
		{
			shop.POST("/search", handler.GetAuthor)

		}

		// 品牌 api/brand/search

		brand := root.Group("/brand")
		{
			brand.POST("/brand", handler.GetAuthor)

		}

		// 工具
		//    爆款探测  api/hot-burst/list
		//    seo 优化  api/seo/hotword/list

	}
}
