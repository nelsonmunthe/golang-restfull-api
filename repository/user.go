package repository

import (
	"anteraja/backend/entity"
	"context"

	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(evoDB *gorm.DB) User {
	return User{
		db: evoDB,
	}
}

type UserInterface interface {
	GetById(ctx context.Context, IDs []uint) ([]entity.UserInterface, error)
}
