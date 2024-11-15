package controllers

import (
	"mini-proyek/config"
	"mini-proyek/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get all Detail Setoran
func GetDetailSetorans(c echo.Context) error {
	var detailSetorans []models.DetailSetoran
	if err := config.DB.Preload("TransaksiSetoran").Preload("JenisSampah").Find(&detailSetorans).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, detailSetorans)
}

// Get Detail Setoran by ID
func GetDetailSetoran(c echo.Context) error {
	id := c.Param("id")
	var detailSetoran models.DetailSetoran
	if err := config.DB.Preload("TransaksiSetoran").Preload("JenisSampah").First(&detailSetoran, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, detailSetoran)
}

// Create new Detail Setoran
func CreateDetailSetoran(c echo.Context) error {
	var detailSetoran models.DetailSetoran
	if err := c.Bind(&detailSetoran); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := config.DB.Create(&detailSetoran).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, detailSetoran)
}

// Update Detail Setoran by ID
func UpdateDetailSetoran(c echo.Context) error {
	id := c.Param("id")
	var detailSetoran models.DetailSetoran
	if err := config.DB.First(&detailSetoran, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	if err := c.Bind(&detailSetoran); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := config.DB.Save(&detailSetoran).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, detailSetoran)
}

// Delete Detail Setoran by ID
func DeleteDetailSetoran(c echo.Context) error {
	id := c.Param("id")
	if err := config.DB.Delete(&models.DetailSetoran{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
