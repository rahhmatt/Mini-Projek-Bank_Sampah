package models

type DetailSetoran struct {
	IDDetailSetoran    uint `gorm:"primaryKey"`
	IDTransaksiSetoran uint
	IDJenisSampah      uint
	Berat              float64
	Subtotal           float64
	TransaksiSetoran   TransaksiSetoran `gorm:"foreignKey:IDTransaksiSetoran"`
	JenisSampah        JenisSampah      `gorm:"foreignKey:IDJenisSampah"`
}

func (DetailSetoran) TableName() string {
	return "detail_setoran"
}
