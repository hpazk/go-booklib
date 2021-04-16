package user

import (
	"net/http"
	"time"

	"github.com/hpazk/go-booklib/helper"
	"github.com/labstack/echo/v4"
)

// TODO 1: payload validator
// TODO 2: existEmail validatation
// TODO 3: createUser service
// TODO 3: authToken service
// TODO 4: response formatter
// TODO 5: api response formatter

// TODO 6: error handling and error-response formatter

func PostUserRegistration(c echo.Context) error {
	req := new(request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.M{"message": "err-badrequest"})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.M{"message": "err-validation"})
	}

	// TODO development: createUser
	newUser := User{
		Name:            req.Name,
		Address:         req.Address,
		Photo:           "",
		Email:           req.Email,
		EmailVerifiedAt: time.Now(),
		Password:        req.Password,
		Role:            "",
	}

	// TODO development: authToken
	authToken := "12345678"

	// TODO development: userData
	userData := userResponseFormatter(newUser, authToken)

	// TODO development: ApiResponseFormatter
	response := helper.ResponseFormatter(http.StatusOK, "success", "user registered", userData)
	return c.JSON(http.StatusOK, response)
}

func PostUserLogin(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "user-login"})
}
