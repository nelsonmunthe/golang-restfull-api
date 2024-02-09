package entity

import (
	"time"
)

type DepositeHeaderEntity struct {
	ID                      uint      `gorm:"primaryKey;type:BIGINT UNSIGNED NOT NULL AUTO_INCREMENT" json:"id"`
	No_transaksi            string    `gorm:"column:no_transaksi;type:VARCHAR  NOT NULL" json:"no_transaksi"`
	Tanggal_transaksi       time.Time `json:"tanggal_transaksi"`
	Last_login              time.Time `json:"last_login"`
	Tanggal_posting         time.Time `json:"tanggal_posting"`
	Status_transaksi        string    `gorm:"column:status_transaksi;type:VARCHAR  NOT NULL" json:"status_transaksi"`
	Status_integrasi        string    `gorm:"column:status_integrasi;type:VARCHAR" json:"status_integrasi"`
	No_transaksi_netsuite   string    `gorm:"column:no_transaksi_netsuite;type:VARCHAR" json:"no_transaksi_netsuite"`
	Subsidiary_id           string    `gorm:"column:subsidiary_id;type:VARCHAR  NOT NULL" json:"subsidiary_id"`
	Subsidiary_name         string    `gorm:"column:subsidiary_name;type:VARCHAR  NOT NULL" json:"subsidiary_name"`
	Location_id             uint      `gorm:"column:location_id;type:BIGINT  NOT NULL" json:"location_id"`
	Location_name           string    `gorm:"column:location_name;type:VARCHAR  NOT NULL" json:"location_name"`
	Area_id                 uint      `gorm:"column:area_id;type:BIGINT  NOT NULL" json:"area_id"`
	Area_name               string    `gorm:"column:area_name;type:VARCHAR  NOT NULL" json:"area_name"`
	Region_id               uint      `gorm:"column:region_id;type:BIGINT  NOT NULL" json:"region_id"`
	Region_name             string    `gorm:"column:region_name;type:VARCHAR  NOT NULL" json:"region_name"`
	Department_id           uint      `gorm:"column:department_id;type:BIGINT  NOT NULL" json:"department_id"`
	Department_name         string    `gorm:"column:department_name;type:VARCHAR(255)  NOT NULL" json:"department_name"`
	Akun_bank_regional      string    `gorm:"column:akun_bank_regional;type:VARCHAR  NOT NULL" json:"akun_bank_regional"`
	Akun_bank_regional_id   string    `gorm:"column:akun_bank_regional_id;type:VARCHAR" json:"akun_bank_regional_id"`
	Akun_bank_regional_nama string    `gorm:"column:akun_bank_regional_nama;type:BIGINT" json:"akun_bank_regional_nama"`
	Total_deposite          float32   `gorm:"column:total_deposite;type:BIGINT" json:"total_deposite"`
	Keterangan              string    `gorm:"column:keterangan;type:VARCHAR  NOT NULL" json:"keterangan"`
	Created_by_id           uint      `gorm:"column:created_by_id;type:BIGINT" json:"created_by_id"`
	Created_by_name         string    `gorm:"column:created_by_name;type:VARCHAR" json:"created_by_name"`
	Created_at              time.Time `json:"created_date"`
	Updated_by              string    `gorm:"column:updated_by;type:VARCHAR" json:"updated_by"`
	Updated_at              time.Time `json:"updated_date"`
	Running_number          uint      `gorm:"column:running_number;type:BIGINT  NOT NULL" json:"running_number"`
	Approve_date            time.Time `json:"approve_date"`
}

func (DepositeHeaderEntity) TableName() string {
	return "trx_deposit_header"
}
