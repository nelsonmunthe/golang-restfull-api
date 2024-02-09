package entity

import (
	"time"
)

type DepositLineEntity struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Role_id    uint      `json:"role_id"`
	Status     bool      `json:"status"`
	Last_login time.Time `json:"last_login"`
	Created_by string    `json:"created_by"`
	Updated_by string    `json:"updated_by"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Viewer     bool      `json:"viewer"`
	UserRole   Role      `gorm:"foreignKey:Role_id"`
}

func (DepositLineEntity) TableName() string {
	return "trx_deposit_header"
}
