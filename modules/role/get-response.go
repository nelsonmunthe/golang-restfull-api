package role

import "anteraja/backend/dto"

type GetResponse struct {
	dto.ResponseMeta
	Data PaginatedGetResponseItems `json:"data"`
}

type PaginatedGetResponseItems struct {
	Pagination dto.PaginationResponse `json:"pagination"`
	Data       []GetResponseItem      `json:"data"`
}

func DefaultGetResponse(
	pagination dto.PaginationResponse,
	data []GetResponseItem,
) GetResponse {
	return GetResponse{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Unit Terbaru MCC",
		},
		Data: PaginatedGetResponseItems{
			Pagination: pagination,
			Data:       data,
		},
	}
}
