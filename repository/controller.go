package repository

import (
	"anteraja/backend/entity"
	"context"

	"gorm.io/gorm"
)

type Controller struct {
	db *gorm.DB
}

func NewController(db *gorm.DB) Controller {
	return Controller{
		db: db,
	}
}

type ControllerInterface interface {
	GetListById(ctx context.Context, userId int) ([]entity.RoleDetailIntity, error)
}

func (repo Controller) GetListById(ctx context.Context, roleId int) ([]entity.RoleDetailIntity, error) {
	var roleDetails []entity.RoleDetailIntity

	err := repo.db.WithContext(ctx).
		Joins("Controller").
		Where("role_id", roleId).
		Where("status", true).
		Find(&roleDetails).
		Error

	return roleDetails, err
}
