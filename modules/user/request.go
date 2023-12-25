package user

type RequestBackfillReparation struct {
	ProductID                           uint   `json:"product_id"  binding:"required"`
	ReparationBudget                    uint   `json:"reparation_budget" binding:"required"`
	OtherBudget                         uint   `json:"other_budget"`
	BankCodeReparation                  string `json:"bank_code_reparation"`
	BankNameReparation                  string `json:"bank_name_reparation" `
	BeneficiaryAccountNumberReparation  string `json:"beneficiary_account_number_reparation"`
	BeneficiaryNameReparation           string `json:"benerficiary_name_reparation"`
	BankCodeOtherBudget                 string `json:"bank_code_other_budget"`
	BankNameOtherBudget                 string `json:"bank_name_other_budget"`
	BeneficiaryAccountNumberOtherBudget string `json:"beneficiary_account_number_other_budget"`
	BeneficiaryNameOtherBudget          string `json:"beneficiary_name_other_budget"`
	MaintenanceInvoice                  string `json:"maintenance_invoice"`
}
