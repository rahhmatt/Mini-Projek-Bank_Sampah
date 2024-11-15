package routes

import (
	"mini-proyek/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {

	// Nasabah routes
	e.GET("/nasabah", controllers.GetNasabahs)
	e.GET("/nasabah/:id", controllers.GetNasabah)
	e.POST("/nasabah", controllers.CreateNasabah)
	e.PUT("/nasabah/:id", controllers.UpdateNasabah)
	e.DELETE("/nasabah/:id", controllers.DeleteNasabah)

	// Transaksi Setoran routes
	e.GET("/transaksi-setoran", controllers.GetTransaksiSetorans)
	e.GET("/transaksi-setoran/:id", controllers.GetTransaksiSetoran)
	e.POST("/transaksi-setoran", controllers.CreateTransaksiSetoran)
	e.PUT("/transaksi-setoran/:id", controllers.UpdateTransaksiSetoran)
	e.DELETE("/transaksi-setoran/:id", controllers.DeleteTransaksiSetoran)

	// Detail Setoran routes
	e.GET("/detail-setoran", controllers.GetDetailSetorans)
	e.GET("/detail-setoran/:id", controllers.GetDetailSetoran)
	e.POST("/detail-setoran", controllers.CreateDetailSetoran)
	e.PUT("/detail-setoran/:id", controllers.UpdateDetailSetoran)
	e.DELETE("/detail-setoran/:id", controllers.DeleteDetailSetoran)

	// Jenis Sampah routes
	e.GET("/jenis-sampah", controllers.GetJenisSampahs)
	e.GET("/jenis-sampah/:id", controllers.GetJenisSampah)
	e.POST("/jenis-sampah", controllers.CreateJenisSampah)
	e.PUT("/jenis-sampah/:id", controllers.UpdateJenisSampah)
	e.DELETE("/jenis-sampah/:id", controllers.DeleteJenisSampah)
}
