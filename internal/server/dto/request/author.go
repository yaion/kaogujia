package request

type AuthorSearchRequest struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"Limit"`
}

type AuthorInfoRequest struct {
	Uid string `json:"uid"`
}

type LiveSearchRequest struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"Limit"`
}

type BrandSearchRequest struct {
}

type ProductSearchRequest struct {
}

type StoreSearchRequest struct {
}

type VideoSearchRequest struct {
}
