package entity

import (
	"time"
)

type UserInterface struct {
	ID                   int       `gorm:"primaryKey" json:"id"`
	UserID               int       `json:"userId"`
	ProductID            int       `json:"productId"`
	BrandID              int       `json:"brandId"`
	ModelID              int       `json:"modelId"`
	Type                 string    `json:"type"`
	Tahun                int       `json:"tahun"`
	PaymentType          string    `json:"paymentType"`
	LeasingID            int       `json:"leasingId"`
	Status               string    `json:"status"`
	ASOID                int       `json:"asoId"`
	SMID                 int       `json:"smId"`
	RegencyID            int       `gorm:"column:regencies_id" json:"regenciesId"`
	AgentID              int       `json:"agentId"`
	StatusProspek        string    `json:"statusProspek"`
	IsRead               int8      `json:"isRead"`
	Name                 string    `json:"name"`
	Phone                string    `json:"phone"`
	CalculatorID         string    `json:"calculatorId"`
	DistrictID           int       `json:"districtId"`
	VillageID            int       `json:"villageId"`
	Address              string    `json:"address"`
	DPPaidAt             time.Time `json:"dpPaidAt"`
	TotalRefund          int       `json:"totalRefund"`
	IncentiveAgentBuyer  int       `json:"incentiveAgentBuyer"`
	IncentiveAgentSeller int       `json:"incentiveAgentSeller"`
	IncentiveAso         int       `json:"incentiveAso"`
	IncentiveMoladin     int       `json:"incentiveMoladin"`
	SubmitterType        string    `json:"submitterType"`
	SubmitterID          int       `json:"submitterId"`
	VerificationStatus   string    `json:"verificationStatus"`
	RedDot               int8      `json:"redDot"`
	FASOID               int       `json:"fasoId"`
	Reason               string    `json:"reason"`
	BankCode             string    `json:"bankCode"`
	BankName             string    `json:"bankName"`
	BankAccountNumber    string    `json:"bankAccountNumber"`
	BankAccountName      string    `json:"bankAccountName"`
	Note                 string    `json:"note"`
	Email                string    `json:"email"`
	IsMoladinInventory   int8      `json:"isMoladinInventory"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}

func (UserInterface) TableName() string {
	return "agent_success_prospek_aso"
}
