package controller

import (
	"net/http"

	model "github.com/CuongPhan99/GolangPostgresql/models"
	"github.com/CuongPhan99/GolangPostgresql/storage"
	"github.com/labstack/echo/v4"
)

// Handler
func GetAccounts(c echo.Context) error {
	accounts, _ := GetRepoAccounts()
	return c.JSON(http.StatusOK, accounts)
}

func GetRepoAccounts() ([]model.Accounts, error) {
	db := storage.GetDBInstance()
	accounts := []model.Accounts{}

	if err := db.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func GetAccountById(c echo.Context) error {
	db := storage.GetDBInstance()
	accounts := []model.Accounts{}

	id := c.Param("id")
	account := db.Where("id = ?", id).First(&accounts)
	if account == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, account)
}

func UpdateLogo(c echo.Context) error {
	db := storage.GetDBInstance()
	accounts := []model.Accounts{}

	id := c.Param("id")
	logo := c.FormValue("logo")
	account := db.Model(&accounts).Where("id = ?", id).Update("logo", logo)

	if account == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, account)
}

func UpdateCompanyName(c echo.Context) error {
	db := storage.GetDBInstance()
	accounts := []model.Accounts{}

	id := c.Param("id")
	company_name := c.FormValue("company_name")
	account := db.Model(&accounts).Where("id = ?", id).Update("company_name", company_name)

	if account == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, account)
}
