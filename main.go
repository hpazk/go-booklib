package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
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
	fmt.Println("starting...")

	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey:  []byte(os.Getenv("JWT_SECRET_KEY")),
	// 	ContextKey:  "user",
	// 	TokenLookup: "header:" + echo.HeaderAuthorization,
	// 	AuthScheme:  "Bearer",
	// 	Claims:      jwt.MapClaims{},
	// }))

	e.Pre(middleware.RemoveTrailingSlash())

	// Static folder images
	e.Static("/", "public")

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helper.M{"message": "success"})
	})

	routes.DefineApiRoutes(e)

	// e.GET("/books", book.GetBooks)
	// e.GET("/books/newest", book.GetNewestBooks)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
