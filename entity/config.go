package entity

import "time"

type ConfigMASOIncentive struct {
	Basic      []ConfigMASOBasicIncentive      `json:"basic"`
	Additional []ConfigMASOAdditionalIncentive `json:"additional"`
}

type ConfigMASOBasicIncentive struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	Label     string    `json:"label"`
	MinPoint  float64   `json:"minPoint"`
	Incentive uint64    `json:"incentive"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type ConfigMASOAdditionalIncentive struct {
	ID                  uint      `gorm:"primaryKey" json:"-"`
	AchievementPoint    float64   `json:"achievementPoint"`
	AdditionalIncentive uint64    `json:"additionalIncentive"`
	CreatedAt           time.Time `json:"-"`
	UpdatedAt           time.Time `json:"-"`
}

type ConfigASOCompensation struct {
	ID              uint      `gorm:"primaryKey" json:"-"`
	Compensation    uint64    `json:"compensation"`
	LowerPriceLimit uint64    `json:"lowerPriceLimit"`
	UpperPriceLimit uint64    `json:"upperPriceLimit"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}

type ConfigMarginPercentagePrice struct {
	ID              uint         `gorm:"primaryKey" json:"-"`
	LowerPriceLimit uint64       `json:"lowerPriceLimit"`
	UpperPriceLimit uint64       `json:"upperPriceLimit"`
	Margin          ConfigMargin `gorm:"foreignKey:PriceRangeID"`
	CreatedAt       time.Time    `json:"-"`
	UpdatedAt       time.Time    `json:"-"`
}

type ConfigMargin struct {
	ID               uint      `gorm:"primaryKey" json:"-"`
	MarginPercentage float64   `json:"marginPercentage"`
	PriceRangeID     uint      `json:"priceRangeId"`
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`
}

type ConfigBuyoutApproval struct {
	ID              uint                `gorm:"primaryKey" json:"-"`
	LowerPriceLimit uint64              `json:"lowerPriceLimit"`
	UpperPriceLimit uint64              `json:"upperPriceLimit"`
	Roles           []ConfigRoleMapping `gorm:"polymorphic:Config;polymorphic_value:buyout_approval,constraint:OnDelete:CASCADE"`
	CreatedAt       time.Time           `json:"-"`
	UpdatedAt       time.Time           `json:"-"`
}

type ConfigReparationBudgetApproval struct {
	ID              uint                `gorm:"primaryKey" json:"-"`
	LowerPriceLimit uint64              `json:"lowerPriceLimit"`
	UpperPriceLimit uint64              `json:"upperPriceLimit"`
	Roles           []ConfigRoleMapping `gorm:"polymorphic:Config;polymorphic_value:reparation_budget_approval,constraint:OnDelete:CASCADE"`
	CreatedAt       time.Time           `json:"-"`
	UpdatedAt       time.Time           `json:"-"`
}

type ConfigPriceChangeApproval struct {
	ID               uint                `gorm:"primaryKey" json:"-"`
	LowerMarginLimit float64             `json:"lowerMarginLimit"`
	UpperMarginLimit float64             `json:"upperMarginLimit"`
	Roles            []ConfigRoleMapping `gorm:"polymorphic:Config;polymorphic_value:price_change_approval,constraint:OnDelete:CASCADE"`
	CreatedAt        time.Time           `json:"-"`
	UpdatedAt        time.Time           `json:"-"`
}

type ConfigBookingFee struct {
	ID              uint      `gorm:"primaryKey" json:"-"`
	Fee             uint64    `json:"fee"`
	LowerPriceLimit uint64    `json:"lowerPriceLimit"`
	UpperPriceLimit uint64    `json:"upperPriceLimit"`
	Repayment       int       `json:"repayment"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}

type ConfigRoleMapping struct {
	ID             uint                 `gorm:"primaryKey" json:"-"`
	RoleID         uint                 `json:"roleId"`
	ConfigID       uint                 `json:"configId"`
	ConfigType     string               `json:"configType"`
	CreatedAt      time.Time            `json:"-"`
	UpdatedAt      time.Time            `json:"-"`
	BuyoutApproval ConfigBuyoutApproval `gorm:"foreignKey:ConfigID"`
}

type ConfigASOBonusPoint struct {
	ID            uint      `json:"-"`
	AsoBonusPoint float64   `json:"asoBonusPoint"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}
