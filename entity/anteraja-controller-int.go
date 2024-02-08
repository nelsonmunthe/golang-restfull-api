package entity

type ControllerEntity struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	Controller_name string `json:"controller_name"`
	Form_id         int    `json:"form_id"`
}

func (ControllerEntity) TableName() string {
	return "ms_controller"
}
