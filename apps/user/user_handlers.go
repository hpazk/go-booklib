package user

import (
	"net/http"
	"strconv"

	"github.com/hpazk/go-booklib/helper"
	"github.com/labstack/echo/v4"
)

// TODO 1: payload validator
// TODO 2: existEmail validatator
// TODO 3: createUser service
// TODO 3: authToken service
// TODO 4: response formatter
// TODO 5: api response formatter

// TODO 6: error handling and error-response formatter

// TODO 7: emailVerification service

func PostUserRegistration(c echo.Context) error {
	req := new(request)
	if err := c.Bind(req); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	// TODO development: createUser
	newUser, _ := signUp(req)

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

func GetUsers(c echo.Context) error {
	s := repo(userStorage{})
	if findByEmail := c.QueryParam("email"); findByEmail != "" {
		s := repo(userStorage{})
		email := c.QueryParam("email")

		response, _ := s.GetByEmail(email)
		return c.JSON(http.StatusOK, response)
	}
	response, _ := s.Fetch()
	return c.JSON(http.StatusOK, response)
}

func GetUser(c echo.Context) error {
	s := repo(userStorage{})
	id, _ := strconv.Atoi(c.Param("id"))
	response, _ := s.GetById(id)
	return c.JSON(http.StatusOK, response)
}

// func GetUserByEmail(c echo.Context) error {
// 	s := repo(userStorage{})
// 	email := c.QueryParam("name")
// 	response, _ := s.GetByEmail(email)
// 	return c.JSON(http.StatusOK, response)
// }

func PutUser(c echo.Context) error {
	s := repo(userStorage{})
	id, _ := strconv.Atoi(c.Param("id"))
	response, _ := s.Update(id)
	return c.JSON(http.StatusOK, response)
}

func DeleteUser(c echo.Context) error {
	s := repo(userStorage{})
	id, _ := strconv.Atoi(c.Param("id"))
	response, _ := s.Delete(id)
	return c.JSON(http.StatusOK, response)
}
