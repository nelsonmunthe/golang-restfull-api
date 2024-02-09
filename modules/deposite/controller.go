package deposit

import (
	"anteraja/backend/dto"
	"anteraja/backend/entity"

	"golang.org/x/net/context"
)

type DepositController struct {
	usecase DepositUsecase
}

type DepositeControllerInterface interface {
	GetList(ctx context.Context, query dto.RequestDeposit, pagination dto.PaginationRequest) ([]entity.DepositeHeaderEntity, error)
}

func (ctrl DepositController) GetList(ctx context.Context, query dto.RequestDeposit, pagination dto.PaginationRequest) (dto.BaseResponseList, error) {
	return ctrl.usecase.GetList(ctx, query, pagination)
}
