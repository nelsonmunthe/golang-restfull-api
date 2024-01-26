package role

import "anteraja/backend/dto"

type QUeryRequest struct {
	dto.PaginationRequest
	Search string `form:"q" json:"q"`
}
