package userv2

import "time"

type RequestUser struct {
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
}

type RequestUserUpdateStatus struct {
	Status bool `json:"status"`
}

type RequestUserUpdateUser struct {
	Password string `json:"password"`
	Role_id  string `json:"role_id"`
	Viewer   bool   `json:"viewer"`
}

type RequestUserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
