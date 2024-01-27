package entity

import (
	"time"
)

type AnterajaUserInt struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Role_id    string    `json:"role_id"`
	Status     bool      `json:"status"`
	Last_login time.Time `json:"last_login"`
	Created_by string    `json:"created_by"`
	Updated_by string    `json:"updated_by"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Viewer     bool      `json:"viewer"`
	Role       Role      `gorm:"foreignKey:Role_id" json:"role"`
}

type AnterajaUserUpdateUserInt struct {
	Password string `json:"password"`
	Role_id  string `json:"role_id"`
	Viewer   bool   `json:"viewer"`
}

type AnterajaUserUpdateStatausInt struct {
	Status bool `json:"status"`
}

type AnterajaUserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AnterajaUserLoginToken struct {
	Token string `json:"token"`
}

func (AnterajaUserInt) TableName() string {
	return "ms_user_revamp"
}
