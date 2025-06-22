package request

type AuthorSearchRequest struct {
	Code string      `json:"author_name"`
	Data interface{} `json:"data"`
}
