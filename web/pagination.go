package web

type PagingRequest struct {
	Offset int `form:"offset" binding:"min=0"`
	Limit  int `form:"limit" binding:"required,min=1,max=100"`
}

type PagingResponse[T any] struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
	Items  []T `json:"items,omitempty"`
}
