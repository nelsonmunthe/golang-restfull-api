package user

import (
	"anteraja/backend/dto"
	"anteraja/backend/repository"
	"context"
)

type UseCase struct {
	userRepo repository.AnterajaUserInterface
}

func (uc UseCase) BackfillReparationProduct(ctx context.Context, request RequestBackfillReparation, userId uint) (dto.ResponseMeta, error) {
	user, err := uc.userRepo.FindById(1)

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

type UseCaseInterface interface {
	BackfillReparationProduct(ctx context.Context, request RequestBackfillReparation, UserID uint) (dto.ResponseMeta, error)
}
