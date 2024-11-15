package models

type JenisSampah struct {
	IDJenisSampah   uint `gorm:"primaryKey"`
	Nama            string
	NamaJenisSampah string
	HargaPerKg      float64
}

func (JenisSampah) TableName() string {
	return "jenis_sampah"
}
