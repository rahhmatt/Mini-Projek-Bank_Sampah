package models

type Nasabah struct {
	ID           int     `gorm:"primaryKey;autoIncrement" json:"ID"`
	Nama         string  `json:"Nama"`
	Alamat       string  `json:"Alamat"`
	NomorTelepon string  `json:"NomorTelepon"`
	Saldo        float64 `json:"Saldo"`
}
