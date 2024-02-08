package entity

import (
	"time"
)

type RoleDetailIntity struct {
	ID            uint             `gorm:"primaryKey" json:"id"`
	Role_id       int              `json:"user_id"`
	Form_id       int              `json:"form_id"`
	Controller_id int              `json:"controller_id"`
	Status        bool             `json:"status"`
	Created_by    string           `json:"created_by"`
	Updated_by    string           `json:"updated_by"`
	Created_at    time.Time        `json:"created_at"`
	Updated_at    time.Time        `json:"updated_at"`
	Controller    ControllerEntity `gorm:"foreignKey:Controller_id"`
}

func (RoleDetailIntity) TableName() string {
	return "ms_roles_detail_revamp"
}
