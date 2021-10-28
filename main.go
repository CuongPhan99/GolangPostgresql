package main

import (
	"net/http"

	"github.com/CuongPhan99/GolangPostgresql/controller"
	"github.com/CuongPhan99/GolangPostgresql/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	storage.NewDb()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	e.GET("/accounts", controller.GetAccounts)

	e.GET("/accounts/:id", controller.GetAccountById)

	e.POST("/accounts/:id/logo", controller.UpdateLogo)

	e.POST("/accounts/:id/company_name", controller.UpdateCompanyName)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
