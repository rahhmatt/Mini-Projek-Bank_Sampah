package models

import (
	"time"
)

type TransaksiSetoran struct {
	IDTransaksiSetoran uint `gorm:"primaryKey"`
	IDNasabah          uint
	TanggalSetoran     time.Time
	TotalBerat         float64
	TotalHarga         float64
	Nasabah            Nasabah `gorm:"foreignKey:IDNasabah"`
}

func (TransaksiSetoran) TableName() string {
	return "transaksi_setoran"
}
