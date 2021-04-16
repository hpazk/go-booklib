package user

import (
	"net/http"

	"github.com/hpazk/go-booklib/helper"
	"github.com/labstack/echo/v4"
)

func PostUserRegistration(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "user-registration"})
}

func PostUserLogin(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "user-login"})
}
