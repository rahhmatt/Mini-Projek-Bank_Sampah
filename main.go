package main

import (
	"mini-proyek/config"
	"mini-proyek/models"
	"mini-proyek/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	config.InitDB()

	config.DB.AutoMigrate(&models.Nasabah{}, &models.TransaksiSetoran{}, &models.DetailSetoran{}, &models.JenisSampah{})

	e := echo.New()

	routes.InitRoutes(e)

	e.Start(":8000")
}
