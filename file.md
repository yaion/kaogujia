## 达人 
- ### 达人列表接口 /api/author/search 
  > 对于视频博主，"video_ratio": 1, "live_ratio": 0
  > 对于直播博主，"live_ratio": 1, "video_ratio": 0
  > urlApi := "https://service.kaogujia.com/api/author/search?limit=50&page=1&sort_field=gmv&sort=0"
  > `{"keyword":"","author_type":0}` // 全部
  > `{"keyword":"","author_type":1}` // 直播
  > `{"keyword":"","author_type":2}`  // 视频 
- ### 直播达人带货榜 接口 api/rank/live/sku/author
> https://service.kaogujia.com/api/rank/live/sku/author?limit=50&page=1&sort_field=gmv&sort=0
- ### 带货达人潜力榜 详情接口 api/rank/author/potential
> https://service.kaogujia.com/api/rank/author/potential?limit=50&page=1&sort_field=potential&sort=0
- ### 达人粉丝榜接口 api/rank/author/fans/increment
> https://service.kaogujia.com/api/rank/author/fans/increment?limit=50&page=1&sort_field=inc_fans_count&sort=0（涨粉）
> https://service.kaogujia.com/api/rank/author/fans/increment?limit=50&page=1&sort_field=inc_fans_count&sort=0(掉粉)  和

- ### 达人详情接口 /darenDetails/darenOverview/{uid}

## 商品 api/sku/search
- ### ?limit=50&page=1&sort_field=sales&sort=0
- ### 实时销量榜
> https://service.kaogujia.com/api/rank/sku/rta?limit=50&page=1&sort_field=h2_sales&sort=0
- ### 商品热销榜
> https://service.kaogujia.com/api/rank/sku/pmt/2?limit=50&page=1&sort_field=sales&sort=0
- ### 直播热推榜
> https://service.kaogujia.com/api/rank/sku/live/popular/2?limit=50&page=1&sort_field=sales&sort=0
- ### 视频热推榜
> https://service.kaogujia.com/api/rank/video/sku/2?limit=50&page=1&sort_field=sales&sort=0

##  直播列表 api/live/search 
> https://service.kaogujia.com/api/live/search?limit=50&page=1&sort_field=gmv&sort=0
- ### 带货小时榜
> https://service.kaogujia.com/api/rank/official/live/sku/hour?limit=50&page=1&sort_field=score&sort=0
- ### 全站小时榜
> https://service.kaogujia.com/api/rank/official/live/hour?limit=50&page=1&sort_field=gap_description&sort=0
- ### 直播大盘 todo

### 视频 /api/video/search
> https://service.kaogujia.com/api/video/search?limit=50&page=1&sort_field=like_count&sort=0
- ### 热门视频榜
> https://service.kaogujia.com/api/rank/video?limit=50&page=1&sort_field=like_count&sort=0
- ### 电商视频榜
> https://service.kaogujia.com/api/rank/productvideo?limit=50&page=1&sort_field=gmv&sort=0
- ### 图文带货榜
> https://service.kaogujia.com/api/rank/productvideo?limit=50&page=1&sort_field=gmv&sort=0
- ### 实时热点榜 todo
> https://service.kaogujia.com/api/rank/sencence/rocketing/20250619/3
> https://service.kaogujia.com/api/rank/sencence/rocketing/20250619/3

## 小店列表 api/shop/search
> https://service.kaogujia.com/api/shop/search?limit=50&page=1&sort_field=gmv&sort=0
- ### 热销小店榜
> https://service.kaogujia.com/api/rank/shop/hot?limit=50&page=1&sort_field=gmv&sort=0
- ### 地区小店榜
> https://service.kaogujia.com/api/rank/shop/area?limit=50&page=1&sort_field=gmv&sort=0

## 品牌 api/brand/search
>https://service.kaogujia.com/api/brand/search?limit=50&page=1&sort_field=gmv&sort=0
- ### 热销品牌榜
> https://service.kaogujia.com/api/rank/brand?limit=50&page=1&sort_field=gmv&sort=0
- ### 品牌声量榜
> https://service.kaogujia.com/api/rank/brandsov?limit=50&page=1&sort_field=expose_count&sort=0

## 工具
- ### 爆款探测
> https://service.kaogujia.com/api/hot-burst/list?limit=25&page=1

- ### 爆款探测
> https://service.kaogujia.com/api/hot-burst/list?limit=25&page=1

- ### 爆款探测
> https://service.kaogujia.com/api/hot-burst/list?limit=25&page=1

- ### seo 优化
> https://service.kaogujia.com/api/seo/hotword/list?limit=50&page=1&sort_field=rise_amount&sort=0
