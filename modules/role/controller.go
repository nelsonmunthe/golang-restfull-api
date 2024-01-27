package role

import (
	"anteraja/backend/dto"

	"golang.org/x/net/context"
)

type RoleController struct {
	usecase RoleUsecase
}

func (role RoleController) GetList(context context.Context, query dto.QUeryRequest) (dto.BaseResponseList, error) {
	return role.usecase.GetList(context, query)
}

func (role RoleController) FindById(context context.Context, userId int) (dto.ResponseMeta, error) {
	return role.usecase.FindById(context, userId)
}
