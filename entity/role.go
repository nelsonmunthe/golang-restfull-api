package entity

import (
	"time"
)

type Role struct {
	ID               int       `gorm:"primaryKey" json:"id"`
	Name             string    `json:"name"`
	Status           bool      `json:"status"`
	Created_by       string    `json:"created_by"`
	Updated_by       string    `json:"updated_by"`
	Created_at       time.Time `json:"created_at"`
	Updated_at       time.Time `json:"updated_at"`
	Default_position string    `json:"default_position"`
}

func (Role) TableName() string {
	return "ms_roles_revamp"
}
