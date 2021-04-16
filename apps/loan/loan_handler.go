package loan

import (
	"net/http"

	"github.com/hpazk/go-booklib/helper"
	"github.com/labstack/echo/v4"
)

func PostLoanBook(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "post-loan"})
}

func GetLoanBook(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "get-loan"})
}

func PutLoanBook(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "put-loan"})
}

func DeleteLoanBook(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "delete-loan"})
}
