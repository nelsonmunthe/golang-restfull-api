package role

import (
	"anteraja/backend/dto"
	"anteraja/backend/repository"
	"anteraja/backend/utils/pagination"
	"context"
)

type RoleUsecase struct {
	roleRepo repository.RoleInterface
}

type RoleUsecaseInterface interface {
	getList(context context.Context) (dto.BaseResponseList, error)
}

func (role RoleUsecase) GetList(context context.Context, query dto.QUeryRequest) (dto.BaseResponseList, error) {
	roles, err := role.roleRepo.GetList(context, query)
	if err != nil {
		return dto.DefaultErrorBaseResponseList(err), err
	}

	count, err := role.roleRepo.CountByCondition(context, query.Search)

	pagination := pagination.Paginate(query.Pagination.Page, query.Pagination.PerPage, int(count))

	if err != nil {
		return dto.DefaultErrorBaseResponseList(err), err
	}

	return dto.BaseResponseList{
		PreviousPage: pagination.PreviousPage,
		CurrentPage:  pagination.CurrentPage,
		NextPage:     pagination.NextPage,
		Total:        int64(pagination.Total),
		PerPage:      pagination.PerPage,
		Data:         roles,
		Success:      true,
		MessageTitle: "",
		Message:      "get Role list successfully",
	}, nil
}

func (role RoleUsecase) FindById(context context.Context, roleId int) (dto.ResponseMeta, error) {
	roleDetail, err := role.roleRepo.FindById(context, roleId)

	if err != nil {
		return defaultErrorResponse(err)
	}

	return dto.ResponseMeta{
		Success:      true,
		MessageTitle: "",
		Message:      "Success get list users",
		ResponseTime: "",
		Data:         roleDetail,
	}, nil

}
