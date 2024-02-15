package services

type Response struct {
	Account_id            string `json:"account_id"`
	Account_subsidiary_id string `json:"account_subsidiary_id"`
	Account_subsidiary    string `json:"account_subsidiary"`
	Account_name          string `json:"account_name"`
	Account_number        string `json:"account_number"`
	Account_type          string `json:"account_type"`
	Is_inactive           bool   `json:"is_inactive"`
	Account_class         string `json:"account_class"`
	Is_pettycash          bool   `json:"is_pettycash"`
}
