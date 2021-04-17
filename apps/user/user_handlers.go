package user

import (
	"fmt"
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

// TODO 8: userHandler

type userHandler struct {
	userServices UserServices
}

func UserHandler(userServices UserServices) *userHandler {
	return &userHandler{userServices}
}

func (h *userHandler) PostUserRegistration(c echo.Context) error {
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

	// TODO development: services
	// TODO development: createUser
	newUser, _ := h.userServices.signUp(req)

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

func (h *userHandler) GetUsers(c echo.Context) error {
	if findByEmail := c.QueryParam("email"); findByEmail != "" {
		email := c.QueryParam("email")

		response, _ := h.userServices.FetchUserByEmail(email)
		return c.JSON(http.StatusOK, response)
	}
	response, _ := h.userServices.FetchUsers()
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	response, _ := h.userServices.FetchUserById(id)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) PutUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := new(updateRequest)

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

	response, _ := h.userServices.UpdateUser(id, req)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.userServices.DeleteUser(id); err != nil {
		return c.JSON(http.StatusOK, helper.M{"message": err})
	}
	message := fmt.Sprintf("user %d was deleted", id)
	return c.JSON(http.StatusOK, helper.M{"message": message})
}
