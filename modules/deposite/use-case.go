package deposit

import (
	"anteraja/backend/dto"
	"anteraja/backend/repository"
	"anteraja/backend/utils/pagination"
	"context"
)

type DepositUsecase struct {
	depositRepo repository.DepositInterface
}

type DepositIUsecaseInterface interface {
	GetList(ctx context.Context, query dto.QUeryRequest, pagination dto.PaginationRequest) (dto.BaseResponseList, error)
	FilterDeposit(ctx context.Context, subsidiary_id string, filter string) (dto.ResponseMeta, error)
}

func (usecase DepositUsecase) GetList(ctx context.Context, query dto.RequestDeposit, page dto.PaginationRequest) (dto.BaseResponseList, error) {
	deposites, err := usecase.depositRepo.GetList(ctx, query, page)
	if err != nil {
		return dto.DefaultErrorBaseResponseList(err), err
	}

	count, err := usecase.depositRepo.CountByCondition(ctx, query)

	pagination := pagination.Paginate(page.Page, page.PerPage, int(count))

	if err != nil {
		return dto.DefaultErrorBaseResponseList(err), err
	}
	return dto.BaseResponseList{
		PreviousPage: pagination.PreviousPage,
		CurrentPage:  pagination.CurrentPage,
		NextPage:     pagination.NextPage,
		Total:        int64(pagination.Total),
		PerPage:      pagination.PerPage,
		Data:         deposites,
		Success:      true,
		MessageTitle: "",
		Message:      "get Deposit list successfully",
	}, nil
}

func (usecase DepositUsecase) FilterDeposit(ctx context.Context, subsidiary_id string, filter string) (dto.ResponseMeta, error) {
	deposites, err := usecase.depositRepo.FilterDeposit(ctx, subsidiary_id, filter)

	if err != nil {
		return defaultErrorResponse(err)
	}

	return dto.ResponseMeta{
		Success:      true,
		MessageTitle: "",
		Message:      "Success get list Filter Deposit",
		ResponseTime: "",
		Data:         deposites,
	}, nil
}
