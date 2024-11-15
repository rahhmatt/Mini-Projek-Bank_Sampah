package controllers

import (
	"mini-proyek/config"
	"mini-proyek/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTransaksiSetorans(c echo.Context) error {
	var transaksiSetorans []models.TransaksiSetoran
	if err := config.DB.Preload("Nasabah").Find(&transaksiSetorans).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transaksiSetorans)
}

func GetTransaksiSetoran(c echo.Context) error {
	id := c.Param("id")
	var transaksiSetoran models.TransaksiSetoran
	if err := config.DB.Preload("Nasabah").First(&transaksiSetoran, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, transaksiSetoran)
}

func CreateTransaksiSetoran(c echo.Context) error {
	var transaksiSetoran models.TransaksiSetoran
	if err := c.Bind(&transaksiSetoran); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := config.DB.Create(&transaksiSetoran).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, transaksiSetoran)
}

func UpdateTransaksiSetoran(c echo.Context) error {
	id := c.Param("id")
	var transaksiSetoran models.TransaksiSetoran
	if err := config.DB.First(&transaksiSetoran, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	if err := c.Bind(&transaksiSetoran); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := config.DB.Save(&transaksiSetoran).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transaksiSetoran)
}

func DeleteTransaksiSetoran(c echo.Context) error {
	id := c.Param("id")
	if err := config.DB.Delete(&models.TransaksiSetoran{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
