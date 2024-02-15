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
	FilterDeposit(ctx context.Context, subsidiary_id string, filter string) (dto.ResponseMeta, error)
	FilterUserByLocation(ctx context.Context, query dto.UserDetailLocationEntity, keyword string) (dto.ResponseMeta, error)
	GetAkuBankRegional(ctx context.Context, params interface{}) (dto.ResponseMeta, error)
}

func (ctrl DepositController) GetList(ctx context.Context, query dto.RequestDeposit, pagination dto.PaginationRequest) (dto.BaseResponseList, error) {
	return ctrl.usecase.GetList(ctx, query, pagination)
}

func (ctrl DepositController) FilterDeposit(ctx context.Context, subsidiary_id string, filter string) (dto.ResponseMeta, error) {
	return ctrl.usecase.FilterDeposit(ctx, subsidiary_id, filter)
}

func (ctrl DepositController) FilterUserByLocation(ctx context.Context, query dto.UserDetailLocationEntity, keyword string) (dto.ResponseMeta, error) {
	return ctrl.usecase.FilterUserByLocation(ctx, query, keyword)
}

func (ctrl DepositController) GetAkuBankRegional(ctx context.Context, params interface{}) (dto.ResponseMeta, error) {
	return ctrl.usecase.GetAkuBankRegional(ctx, params)
}
