package repository

import (
	"anteraja/backend/entity"

	"gorm.io/gorm"
)

type AnterajaUser struct {
	db *gorm.DB
}

func NewAnterajaUser(evoDB *gorm.DB) AnterajaUser {
	return AnterajaUser{
		db: evoDB,
	}
}

func (repo AnterajaUser) FindById(id interface{}) (entity.AnterajaUserInt, error) {
	var user entity.AnterajaUserInt
	err := repo.db.First(&user, id).Error
	return user, err
}

type AnterajaUserInterface interface {
	FindById(id interface{}) (entity.AnterajaUserInt, error)
}
