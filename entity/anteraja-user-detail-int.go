package entity

import (
	"time"
)

type UserDetailEntity struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	User_id           string    `json:"user_id"`
	Location_id       int       `json:"location_id"`
	Location_name     string    `json:"location_name"`
	Subsidiary_id     string    `json:"subsidiary_id"`
	Subsidiary_name   string    `json:"subsidiary_name"`
	Area_id           int       `json:"area_id"`
	Area_name         string    `json:"area_name"`
	Region_id         int       `json:"region_id"`
	Region_name       string    `json:"region_name"`
	Layer_1           int       `json:"layer_1"`
	Max_limit_layer_1 int       `json:"max_limit_layer_1"`
	Layer_2           int       `json:"layer_2"`
	Max_limit_layer_2 int       `json:"max_limit_layer_2"`
	Layer_3           int       `json:"layer_3"`
	Max_limit_layer_3 int       `json:"max_limit_layer_3"`
	Layer_4           int       `json:"layer_4"`
	Max_limit_layer_4 int       `json:"max_limit_layer_4"`
	Status            bool      `json:"status"`
	Created_by        string    `json:"created_by"`
	Updated_by        string    `json:"updated_by"`
	Created_at        time.Time `json:"created_at"`
	Updated_at        time.Time `json:"updated_at"`
}

func (UserDetailEntity) TableName() string {
	return "ms_user_detail_revamp"
}

type UserDetailLocationEntity struct {
	Location_id   int    `json:"location_id"`
	Location_name string `json:"location_name"`
	Area_id       int    `json:"area_id"`
	Area_name     string `json:"area_name"`
	Region_id     int    `json:"region_id"`
	Region_name   string `json:"region_name"`
}

func (UserDetailLocationEntity) TableName() string {
	return "ms_user_detail_revamp"
}
