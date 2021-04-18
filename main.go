package main

import (
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/hpazk/go-booklib/database"
	"github.com/hpazk/go-booklib/helper"
	"github.com/hpazk/go-booklib/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()

	// Custom Validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Trailing slash
	e.Pre(middleware.RemoveTrailingSlash())

	// Static folder images
	e.Static("/", "public")

	// Main root
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helper.M{"message": "success"})
	})

	// Db
	db := database.GetDbInstance()
	dbMigration := database.GetMigrations(db)
	err := dbMigration.Migrate()
	if err == nil {
		print("Migrations did run successfully")
	} else {
		print("migrations failed.", err)
	}
	// Routes
	routes.DefineApiRoutes(e)

	// e.GET("/books", book.GetBooks)
	// e.GET("/books/newest", book.GetNewestBooks)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
