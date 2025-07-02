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
			author.POST("/search", handler.GetAuthorList)

		}

		// 商品 api/sku/search
		sku := root.Group("/sku")
		{
			sku.POST("/search", handler.GetAuthorInfo)

		}

		// 直播列表 api/live/search

		live := root.Group("/live")
		{
			live.POST("/search", handler.GetLiveList)

		}

		// 视频 /api/video/search

		video := root.Group("/video")
		{
			video.POST("/search", handler.GetVideoList)

		}

		// 小店列表 api/shop/search

		shop := root.Group("/shop")
		{
			shop.POST("/search", handler.GetStoreList)

		}

		// 品牌 api/brand/search

		brand := root.Group("/brand")
		{
			brand.POST("/search", handler.GetBrandList)

		}

		// 商品 api/product/search

		product := root.Group("/product")
		{
			product.POST("/search", handler.GetProductList)

		}

		// 工具
		//    爆款探测  api/hot-burst/list
		//    seo 优化  api/seo/hotword/list

	}
}
