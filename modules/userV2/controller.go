package userv2

import (
	"anteraja/backend/dto"
	"context"
)

type UserController struct {
	useCase UserV2InterfaceUsecase
}

func (c UserController) FindByIdV2(context context.Context, userId uint) (dto.ResponseMeta, error) {
	return c.useCase.FindByIdV2(context, userId)
}

func (c UserController) GetListUser(context context.Context) (dto.ResponseMeta, error) {
	return c.useCase.GetListUser(context)
}

func (c UserController) DeleteUser(context context.Context, userId uint) (dto.ResponseMeta, error) {
	return c.useCase.DeleteUser(context, userId)
}

func (c UserController) UpdateUser(context context.Context, userId uint, request RequestUserUpdateUser) (dto.ResponseMeta, error) {
	return c.useCase.UpdateUser(context, userId, request)
}

func (c UserController) CreateUser(context context.Context, request RequestUser) (dto.ResponseMeta, error) {
	return c.useCase.CreateUser(context, request)
}

func (c UserController) changeStatus(context context.Context, userId uint, request RequestUserUpdateStatus) (dto.ResponseMeta, error) {
	return c.useCase.changeStatus(context, userId, request)
}

func (c UserController) login(context context.Context, request RequestUserLogin) (dto.ResponseMeta, error) {
	return c.useCase.login(context, request)
}
