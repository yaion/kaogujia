package service

import (
	"fmt"
	"kaogujia/internal/server/dao/mongodb"
	"kaogujia/pkg/config"
	"kaogujia/pkg/middleware"
	"log"
	"reflect"
	"testing"
)

func TestCreateAuthors(t *testing.T) {
	// 1. 加载配置
	if err := config.Load("D:/codes/go/kaogujia/configs/app.yaml"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. 初始化中间件 (传递配置)
	cfg := config.Get()
	fmt.Println(cfg)
	if err := middleware.InitAll(); err != nil {
		log.Fatalf("Middleware init failed: %v", err)
	}
	type args struct {
		authors []*mongodb.Author
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				authors: []*mongodb.Author{
					{
						Aup:             "1",
						Avatar:          "1",
						AvgLikeCount:    "1",
						AvgLiveGmv:      "1",
						AvgLiveSales:    "1",
						AvgPlayCount:    "1",
						AvgStayDuration: 1,
						AvgThrough:      1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateAuthors(tt.args.authors)
		})
	}
}

func TestGetAuthorList(t *testing.T) {
	// 1. 加载配置
	if err := config.Load("D:/codes/go/kaogujia/configs/app.yaml"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. 初始化中间件 (传递配置)
	cfg := config.Get()
	fmt.Println(cfg)
	if err := middleware.InitAll(); err != nil {
		log.Fatalf("Middleware init failed: %v", err)
	}
	type args struct {
		authors []*mongodb.Author
	}
	tests := []struct {
		name    string
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "test",
			want: map[string]interface{}{
				"total": 0,
				"authors": []interface{}{
					map[string]interface{}{
						"aup":            "1",
						"avatar":         "1",
						"avg_like_count": "1",
						"avg_live_gmv":   "1",
						"avg_live_sales": "1",
						"avg_play_count": "1",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAuthorList()
			fmt.Println(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAuthorList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAuthorList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
