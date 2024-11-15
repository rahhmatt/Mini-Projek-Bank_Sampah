package controllers

import (
	"mini-proyek/config"
	"mini-proyek/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetNasabahs(c echo.Context) error {
	var nasabahs []models.Nasabah
	if err := config.DB.Find(&nasabahs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nasabahs)
}

func GetNasabah(c echo.Context) error {
	id := c.Param("id")
	var nasabah models.Nasabah
	if err := config.DB.First(&nasabah, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, nasabah)
}

func CreateNasabah(c echo.Context) error {
	var nasabah models.Nasabah
	if err := c.Bind(&nasabah); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := config.DB.Create(&nasabah).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, nasabah)
}

func UpdateNasabah(c echo.Context) error {
	id := c.Param("id")
	var nasabah models.Nasabah
	if err := config.DB.First(&nasabah, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	if err := c.Bind(&nasabah); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := config.DB.Save(&nasabah).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nasabah)
}

func DeleteNasabah(c echo.Context) error {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Nasabah{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
