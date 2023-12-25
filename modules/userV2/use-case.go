package userv2

import (
	"anteraja/backend/dto"
	"anteraja/backend/entity"
	"anteraja/backend/middleware"
	"anteraja/backend/repository"
	bcryptpassword "anteraja/backend/utils/bcryptPassword"
	"context"
	"fmt"
)

type UserV2Usecase struct {
	userRepoV2 repository.AnterajaUserInterfaceV2
}

type UserV2InterfaceUsecase interface {
	FindByIdV2(context context.Context, userId uint) (dto.ResponseMeta, error)
	GetListUser(context context.Context) (dto.ResponseMeta, error)
	DeleteUser(context context.Context, userId uint) (dto.ResponseMeta, error)
	UpdateUser(context context.Context, userId uint, request RequestUserUpdateUser) (dto.ResponseMeta, error)
	CreateUser(context context.Context, request RequestUser) (dto.ResponseMeta, error)
	changeStatus(context context.Context, userId uint, request RequestUserUpdateStatus) (dto.ResponseMeta, error)
	login(context context.Context, request RequestUserLogin) (dto.ResponseMeta, error)
}

func (uc UserV2Usecase) FindByIdV2(context context.Context, userId uint) (dto.ResponseMeta, error) {
	user, err := uc.userRepoV2.FindById(userId)

	if err != nil {
		return defaultErrorResponse(err)
	}

	return dto.ResponseMeta{
		Success:      true,
		MessageTitle: "",
		Message:      "Success get user detail",
		ResponseTime: "",
		Data:         user,
	}, nil
}

func (uc UserV2Usecase) GetListUser(context context.Context) (dto.ResponseMeta, error) {
	users, err := uc.userRepoV2.GetListUser()
	for i := range users {
		fmt.Println("index", users[i].Username)
	}

	if err != nil {
		return defaultErrorResponse(err)
	}

	return dto.ResponseMeta{
		Success:      true,
		MessageTitle: "",
		Message:      "Success get list users",
		ResponseTime: "",
		Data:         users,
	}, nil
}

func (uc UserV2Usecase) DeleteUser(context context.Context, userId uint) (dto.ResponseMeta, error) {
	user, err := uc.userRepoV2.DeleteUser(userId)

	if err != nil {
		return defaultErrorResponse(err)
	}

	return dto.ResponseMeta{
		Success:      true,
		MessageTitle: "",
		Message:      "Success delete user",
		ResponseTime: "",
		Data:         user,
	}, nil
}

func (uc UserV2Usecase) UpdateUser(context context.Context, userId uint, request RequestUserUpdateUser) (dto.ResponseMeta, error) {
	newUpdateUser := &entity.AnterajaUserUpdateUserInt{
		Password: request.Password,
		Role_id:  request.Role_id,
		Viewer:   request.Viewer,
	}

	err := uc.userRepoV2.UpdateUser(context, userId, *newUpdateUser)
	if err != nil {
		return defaultErrorResponse((err))
	}

	return dto.ResponseMeta{
		Success:      true,
		MessageTitle: "",
		Message:      "Success update user",
		ResponseTime: "",
		Data:         newUpdateUser,
	}, nil
}

func (uc UserV2Usecase) CreateUser(context context.Context, request RequestUser) (dto.ResponseMeta, error) {
	hash, _ := bcryptpassword.HashPassword(request.Password)

	newUser := &entity.AnterajaUserInt{
		ID:         request.ID,
		Username:   request.Username,
		Password:   hash,
		Role_id:    request.Role_id,
		Status:     request.Status,
		Last_login: request.Last_login,
		Created_by: request.Created_by,
		Updated_by: request.Updated_by,
		Created_at: request.Created_at,
		Updated_at: request.Updated_at,
		Viewer:     request.Viewer,
	}

	err := uc.userRepoV2.CreateUser(context, *newUser)
	if err != nil {
		return defaultErrorResponse((err))
	}

	return dto.ResponseMeta{
		Success:      true,
		MessageTitle: "",
		Message:      "Success Create user",
		ResponseTime: "",
		Data:         newUser,
	}, nil
}

func (uc UserV2Usecase) changeStatus(context context.Context, userId uint, request RequestUserUpdateStatus) (dto.ResponseMeta, error) {
	user, err := uc.userRepoV2.FindById(userId)
	if err != nil {
		return defaultErrorResponse((err))
	}

	status := &entity.AnterajaUserUpdateStatausInt{
		Status: request.Status,
	}

	err = uc.userRepoV2.ChangeStatus(context, user.ID, *status)

	if err != nil {
		return defaultErrorResponse((err))
	}

	return dto.ResponseMeta{
		Success:      true,
		MessageTitle: "",
		Message:      "Success update user",
		ResponseTime: "",
		Data:         user,
	}, nil
}

func (uc UserV2Usecase) login(context context.Context, request RequestUserLogin) (dto.ResponseMeta, error) {

	user, err := uc.userRepoV2.FindByUsername(request.Username)

	if err != nil {
		return defaultErrorResponse(err)
	}

	match := bcryptpassword.CheckPasswordHash(request.Password, user.Password)
	if match == false {
		return defaultErrorResponse(err)
	}

	token, err := middleware.GenerateJwtToken(request.Username)
	if err != nil {
		return defaultErrorResponse(err)
	}
	return dto.ResponseMeta{
		Success:      true,
		MessageTitle: "",
		Message:      "Success Login user",
		ResponseTime: "",
		Data:         token,
	}, nil

}
