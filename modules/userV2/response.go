package userv2

import "anteraja/backend/dto"

func defaultErrorResponse(err error) (dto.ResponseMeta, error) {
	return dto.ResponseMeta{
		Success:      false,
		MessageTitle: "Oops, something went wrong.",
		Message:      err.Error(),
		ResponseTime: "",
	}, err
}

type LoginResponse struct {
	Token       string       `json:"token"`
	User_id     int          `json:"user_id"`
	Username    string       `json:"username"`
	Role_id     int          `json:"role_id"`
	Role        string       `json:"role"`
	Controllers []string     `json:"controllers"`
	Subsidiarys []Subsidiary `json:"subsidiarys"`
}

type Subsidiary struct {
	Subsidiary_id   string `json:"subsidiary_id"`
	Subsidiary_name string `json:"subsidiary_name"`
}

type TokenPayload struct {
	Username    string `json:"username"`
	ID          int    `json:"id"`
	Position_id string `json:"position_id"`
	Role_id     int    `json:"role_id"`
	Viewer      bool   `json:"viewer"`
	Areas       []uint `json:"areas"`
}
