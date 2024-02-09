package dto

type RequestDeposit struct {
	No_transaksi            string `form:"no_transaksi" json:"no_transaksi"`
	Tanggal_transaksi_start string `form:"tanggal_transaksi_start" json:"tanggal_transaksi_start"`
	Location_name           string `form:"location_name" json:"location_name"`
	Region_name             string `form:"region_name" json:"region_name"`
	Status_integrasi        string `form:"status_integrasi" json:"status_integrasi"`
	Status_transaksi        string `form:"status_transaksi" json:"status_transaksi"`
	Tanggal_posting         string `form:"tanggal_posting" json:"tanggal_posting"`
	Akun_bank_regional      string `form:"akun_bank_regional" json:"akun_bank_regional"`
	Subsidiary_id           string `form:"subsidiary_id" json:"subsidiary_id"`
}
