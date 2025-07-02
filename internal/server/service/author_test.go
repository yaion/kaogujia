package service

import (
	"encoding/json"
	"fmt"
	"kaogujia/internal/server/dao/mongodb"
	"kaogujia/pkg/config"
	"kaogujia/pkg/middleware"
	"log"
	"os"
	"testing"
)

type AuthorResult struct {
	IsAuthority bool              `json:"is_authority"`
	Items       []*mongodb.Author `json:"items"`
	Pagination  Pagination        `json:"pagination"`
	Sort        Sort              `json:"sort"`
}

type BrandResult struct {
	IsAuthority bool             `json:"is_authority"`
	Items       []*mongodb.Brand `json:"items"`
	Pagination  Pagination       `json:"pagination"`
	Sort        Sort             `json:"sort"`
}

type ProductResult struct {
	IsAuthority bool               `json:"is_authority"`
	Items       []*mongodb.Product `json:"items"`
	Pagination  Pagination         `json:"pagination"`
	Sort        Sort               `json:"sort"`
}

type liveResult struct {
	IsAuthority bool            `json:"is_authority"`
	Items       []*mongodb.Live `json:"items"`
	Pagination  Pagination      `json:"pagination"`
	Sort        Sort            `json:"sort"`
}

type Pagination struct {
	TotalCount int64 `json:"total_count"`
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
}

type Sort struct {
	SortField string `json:"sort_field"`
	Sort      int    `json:"sort"`
}

func getFile(fiel string) []byte {
	filePath := fmt.Sprintf("E:/code/go/kaogujia/%s", fiel)

	// 读取文件内容
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return nil
	}

	return data
}

func TestCreateAuthors(t *testing.T) {
	// 1. 加载配置
	if err := config.Load("E:/code/go/kaogujia/configs/app.yaml"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. 初始化中间件 (传递配置)
	cfg := config.Get()
	fmt.Println(cfg)
	if err := middleware.InitAll(); err != nil {
		log.Fatalf("Middleware init failed: %v", err)
	}

	// 读取文件内容
	data := getFile("author.json")
	//fmt.Println(string(data))
	authorResult := new(AuthorResult)
	err := json.Unmarshal(data, authorResult)
	if err != nil {
		fmt.Println("json.Unmarshal failed:", err)
		return
	}

	fmt.Println(authorResult)

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
				authors: authorResult.Items,
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
	if err := config.Load("E:/code/go/kaogujia/configs/app.yaml"); err != nil {
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
			/*got, err := GetAuthorList()
			fmt.Println(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAuthorList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAuthorList() got = %v, want %v", got, tt.want)
			}*/
		})
	}
}
