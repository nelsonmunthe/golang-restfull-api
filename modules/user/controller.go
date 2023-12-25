package user

import (
	"anteraja/backend/dto"
	"context"
)

type Controller struct {
	useCase UseCaseInterface
}

func (ctrl Controller) backfillReparationProduct(ctx context.Context, request RequestBackfillReparation, UserID uint) (dto.ResponseMeta, error) {
	return ctrl.useCase.BackfillReparationProduct(ctx, request, UserID)
}
