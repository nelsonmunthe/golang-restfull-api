package deposit

import "anteraja/backend/dto"

func defaultErrorResponse(err error) (dto.ResponseMeta, error) {
	return dto.ResponseMeta{
		Success:      false,
		MessageTitle: "Oops, something went wrong.",
		Message:      err.Error(),
		ResponseTime: "",
	}, err
}
