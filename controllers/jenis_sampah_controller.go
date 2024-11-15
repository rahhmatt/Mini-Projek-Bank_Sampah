package controllers

import (
	"mini-proyek/config"
	"mini-proyek/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetJenisSampahs(c echo.Context) error {
	var jenisSampahs []models.JenisSampah
	if err := config.DB.Find(&jenisSampahs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, jenisSampahs)
}

func GetJenisSampah(c echo.Context) error {
	id := c.Param("id")
	var jenisSampah models.JenisSampah
	if err := config.DB.First(&jenisSampah, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, jenisSampah)
}

func CreateJenisSampah(c echo.Context) error {
	var jenisSampah models.JenisSampah
	if err := c.Bind(&jenisSampah); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := config.DB.Create(&jenisSampah).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, jenisSampah)
}

func UpdateJenisSampah(c echo.Context) error {
	id := c.Param("id")
	var jenisSampah models.JenisSampah
	if err := config.DB.First(&jenisSampah, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	if err := c.Bind(&jenisSampah); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := config.DB.Save(&jenisSampah).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, jenisSampah)
}

func DeleteJenisSampah(c echo.Context) error {
	id := c.Param("id")
	if err := config.DB.Delete(&models.JenisSampah{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
