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
	Area_name               string `form:"area_name" json:"area_name"`
	Subsidiary_id           string `form:"subsidiary_id" json:"subsidiary_id"`
}

type RequestFilterDeposit struct {
	No_transaksi            bool   `form:"no_transaksi" json:"no_transaksi"`
	Tanggal_transaksi_start bool   `form:"tanggal_transaksi_start" json:"tanggal_transaksi_start"`
	Location_name           bool   `form:"location_name" json:"location_name"`
	Region_name             bool   `form:"region_name" json:"region_name"`
	Status_integrasi        bool   `form:"status_integrasi" json:"status_integrasi"`
	Status_transaksi        bool   `form:"status_transaksi" json:"status_transaksi"`
	Tanggal_posting         bool   `form:"tanggal_posting" json:"tanggal_posting"`
	Akun_bank_regional      bool   `form:"akun_bank_regional" json:"akun_bank_regional"`
	Area_name               bool   `form:"area_name" json:"area_name"`
	Subsidiary_id           string `form:"subsidiary_id" json:"subsidiary_id"`
}
