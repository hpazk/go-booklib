package main

import (
	"fmt"
	"net/http"

	"github.com/hpazk/go-booklib/apps/book"
	"github.com/hpazk/go-booklib/helper"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("starting...")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helper.M{"message": "success"})
	})

	e.GET("/books", book.GetBookCatalog)
	e.GET("/books/newest", book.GetBookCatalog)

	e.Logger.Fatal(e.Start(":8080"))
}
