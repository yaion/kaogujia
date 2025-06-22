package response

type AuthorSearchResponse struct {
	Code string      `json:"author_name"`
	Data interface{} `json:"data"`
}
