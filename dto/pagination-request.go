package dto

type PaginationRequest struct {
	Page    int `form:"page" json:"p"`
	PerPage int `form:"perPage" json:"pp"`
}
