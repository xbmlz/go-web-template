package dto

type PageRequest struct {
	Page int `json:"page" form:"page" validate:"required,number"`
	Size int `json:"size" form:"size" validate:"required,number"`
}

type PageResponse[T any] struct {
	Page  int   `json:"page"`
	Size  int   `json:"size"`
	Total int64 `json:"total"`
	List  []T   `json:"list"`
}
