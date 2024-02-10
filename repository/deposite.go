package repository

import (
	"anteraja/backend/dto"
	"anteraja/backend/entity"
	"anteraja/backend/utils/pagination"
	"context"

	"gorm.io/gorm"
)

type Deposit struct {
	db *gorm.DB
}

func NewDeposit(db *gorm.DB) Deposit {
	return Deposit{
		db: db,
	}
}

type DepositInterface interface {
	GetList(ctx context.Context, query dto.RequestDeposit, page dto.PaginationRequest) ([]entity.DepositeHeaderEntity, error)
	CountByCondition(ctx context.Context, query dto.RequestDeposit) (uint, error)
	FilterDeposit(ctx context.Context, subsidiary_id string, filter string) ([]interface{}, error)
}

func (deposit Deposit) GetList(ctx context.Context, query dto.RequestDeposit, page dto.PaginationRequest) ([]entity.DepositeHeaderEntity, error) {
	var deposits []entity.DepositeHeaderEntity
	offset, limit := pagination.CountLimitAndOffset(page.Page, page.PerPage)

	err := deposit.db.
		Where(query).
		Offset(offset).
		Limit(limit).
		Find(&deposits).
		Error
	return deposits, err
}

func (deposit Deposit) CountByCondition(ctx context.Context, query dto.RequestDeposit) (uint, error) {
	var count int64
	countQuery := deposit.db.WithContext(ctx).Model(&entity.DepositeHeaderEntity{})
	countQuery = deposit.db.WithContext(ctx).
		Model(&entity.DepositeHeaderEntity{}).
		Where(query)
	err := countQuery.Count(&count).Error
	return uint(count), err
}

func (deposit Deposit) FilterDeposit(ctx context.Context, subsidiary_id string, filter string) ([]interface{}, error) {
	var deposits []interface{}

	err := deposit.db.WithContext(ctx).
		Table("trx_deposit_header").
		Select(filter).Where("subsidiary_id = ? ", subsidiary_id).
		Group(filter).
		Find(&deposits).Error
	return deposits, err
}
