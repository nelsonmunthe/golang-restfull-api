package repository

import (
	"anteraja/backend/entity"
	"context"

	"gorm.io/gorm"
)

type UserDetail struct {
	db *gorm.DB
}

func NewUserDetail(db *gorm.DB) UserDetail {
	return UserDetail{
		db: db,
	}
}

type UserDetailInterface interface {
	GetListById(ctx context.Context, userId int) ([]entity.UserDetailEntity, error)
	FindUserDetail(ctx context.Context, userId int, subSidiaryId string) (entity.UserDetailEntity, error)
}

func (repo UserDetail) GetListById(ctx context.Context, userId int) ([]entity.UserDetailEntity, error) {
	var userDetails []entity.UserDetailEntity

	err := repo.db.WithContext(ctx).
		Select("area_id").
		Where("user_id", userId).
		Where("status", true).
		Group("area_id").
		Find(&userDetails).
		Error

	return userDetails, err
}

func (repo UserDetail) FindUserDetail(ctx context.Context, userId int, subSidiaryId string) (entity.UserDetailEntity, error) {
	var userDetail entity.UserDetailEntity

	err := repo.db.WithContext(ctx).
		Where("user_id", userId).
		Where("subsidiary_id", subSidiaryId).
		Where("status", true).
		Find(&userDetail).
		Error

	return userDetail, err
}
