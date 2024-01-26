package repository

import (
	"anteraja/backend/dto"
	"anteraja/backend/entity"
	"anteraja/backend/utils/pagination"
	"context"

	"gorm.io/gorm"
)

type Role struct {
	db *gorm.DB
}

func NewRole(db *gorm.DB) Role {
	return Role{
		db: db,
	}
}

type RoleInterface interface {
	GetList(ctx context.Context, query dto.QUeryRequest) ([]entity.Role, error)
	CountByCondition(ctx context.Context, q string) (uint, error)
}

func (role Role) GetList(ctx context.Context, query dto.QUeryRequest) ([]entity.Role, error) {
	var roles []entity.Role

	offset, limit := pagination.CountLimitAndOffset(query.Pagination.Page, query.Pagination.PerPage)
	if query.Search != "" {
		err := role.db.
			Where("name = ?", query.Search).
			Offset(offset).
			Limit(limit).
			Find(&roles).
			Order("`id` desc").Error
		return roles, err
	}
	err := role.db.
		Offset(offset).
		Limit(limit).
		Find(&roles).
		Order("`id` desc").Error
	return roles, err
}

func (role Role) CountByCondition(ctx context.Context, q string) (uint, error) {
	var count int64
	countQuery := role.db.WithContext(ctx).Model(&entity.Role{})
	if q != "" {
		countQuery = role.db.WithContext(ctx).
			Model(&entity.Role{}).
			Where("name = ?", q)
	}
	err := countQuery.Count(&count).Error
	return uint(count), err
}
