package controllers

import (
	"api-echo-go/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchAllDosen(c echo.Context) error {
	result, err := models.FetchAllDosen()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
