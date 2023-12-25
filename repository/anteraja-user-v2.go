package repository

import (
	"anteraja/backend/entity"
	"context"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AnterajaUserV2 struct {
	db *gorm.DB
}

func NewAnterajaUserV2(evoDB *gorm.DB) AnterajaUserV2 {
	return AnterajaUserV2{
		db: evoDB,
	}
}

func (repo AnterajaUserV2) FindById(id interface{}) (entity.AnterajaUserInt, error) {
	var user entity.AnterajaUserInt
	err := repo.db.First(&user, id).Error
	return user, err
}

func (repo AnterajaUserV2) GetListUser() ([]entity.AnterajaUserInt, error) {
	var user []entity.AnterajaUserInt
	err := repo.db.Find(&user).Error
	return user, err
}

func (repo AnterajaUserV2) DeleteUser(userId interface{}) (entity.AnterajaUserInt, error) {
	var user entity.AnterajaUserInt
	err := repo.db.Delete(&user, userId).Error
	return user, err
}

func (repo AnterajaUserV2) UpdateUser(context context.Context, userId interface{}, updateuser entity.AnterajaUserUpdateUserInt) error {

	var user *entity.AnterajaUserInt
	err := repo.db.Model(user).WithContext(context).
		Where("id = ?", userId).
		Update("Password", updateuser.Password).
		Error
	return err
}

func (repo AnterajaUserV2) CreateUser(context context.Context, newUser entity.AnterajaUserInt) error {
	err := repo.db.WithContext(context).Save(&newUser).Error
	return err
}

func (repo AnterajaUserV2) ChangeStatus(context context.Context, userId interface{}, request entity.AnterajaUserUpdateStatausInt) error {
	var user *entity.AnterajaUserInt
	err := repo.db.Model(user).WithContext(context).
		Where("id = ?", userId).
		Update("status", request.Status).
		Error
	return err

}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (repo AnterajaUserV2) FindByUsername(username interface{}) (entity.AnterajaUserInt, error) {
	var user entity.AnterajaUserInt
	err := repo.db.Model(&user).Where("username = ?", username).Find(&user).Error
	return user, err
}

type AnterajaUserInterfaceV2 interface {
	FindById(id interface{}) (entity.AnterajaUserInt, error)
	GetListUser() ([]entity.AnterajaUserInt, error)
	DeleteUser(userId interface{}) (entity.AnterajaUserInt, error)
	UpdateUser(context context.Context, userId interface{}, request entity.AnterajaUserUpdateUserInt) error
	CreateUser(context context.Context, newUser entity.AnterajaUserInt) error
	ChangeStatus(context context.Context, userId interface{}, request entity.AnterajaUserUpdateStatausInt) error
	FindByUsername(username interface{}) (entity.AnterajaUserInt, error)
}
