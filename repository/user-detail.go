package repository

import (
	"anteraja/backend/dto"
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
	FilterUserByLocation(ctx context.Context, query dto.UserDetailLocationEntity, keyword string) ([]entity.UserDetailLocationEntity, error)
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

func (repo UserDetail) FilterUserByLocation(ctx context.Context, query dto.UserDetailLocationEntity, keyword string) ([]entity.UserDetailLocationEntity, error) {
	var userDetails []entity.UserDetailLocationEntity

	var err error

	if keyword == "location" {
		err = repo.db.WithContext(ctx).
			Select("location_id", "location_name", "area_id", "area_name", "region_id", "region_name").
			Where(query).
			Group("location_id, location_name, area_id, area_name, region_id, region_name").
			Find(&userDetails).
			Error

	} else if keyword == "area" {
		err = repo.db.WithContext(ctx).
			Select("area_id", "area_name", "region_id", "region_name").
			Where(query).
			Group("area_id, area_name, region_id, region_name").
			Find(&userDetails).
			Error

	} else if keyword == "region" {
		err = repo.db.WithContext(ctx).
			Select("region_id", "region_name").
			Where(query).
			Group("region_id, region_name").
			Find(&userDetails).
			Error
	}

	return userDetails, err
}
