package book

import (
	"fmt"
	"net/http"

	"github.com/hpazk/go-booklib/helper"
	"github.com/labstack/echo/v4"
)

func GetBookCatalog(c echo.Context) error {
	// TODO 1: get books
	// TODO 1: get books by category
	category := c.QueryParam("category")
	fmt.Println(category)

	return c.JSON(http.StatusOK, helper.M{"message": "book-catalog"})
}
func GetBookNewest(c echo.Context) error {
	// TODO 2: get books by newest year

	return c.JSON(http.StatusOK, helper.M{"message": "book-newest"})
}
