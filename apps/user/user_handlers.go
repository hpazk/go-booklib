package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/hpazk/go-booklib/helper"
	"github.com/labstack/echo/v4"
)

// TODO 1: payload validator: OK
// TODO 2: existEmail validatator
// TODO 3: createUser service: OK
// TODO 3: authToken service
// TODO 4: response formatter: OK
// TODO 5: api response formatter: OK
// TODO 6: error handling and error-response formatter: OK
// TODO 7: emailVerification service
// TODO 8: userHandler: OK
// TODO 9: handler login
// TODO 10: handler logout

type userHandler struct {
	userServices UserServices
}

func UserHandler(userServices UserServices) *userHandler {
	return &userHandler{userServices}
}

func (h *userHandler) PostUserRegistration(c echo.Context) error {
	req := new(request)
	if err := c.Bind(req); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", "invalid request", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	newUser, _ := h.userServices.signUp(req)

	// TODO development: authToken
	authToken := "12345678"

	// TODO development: userData
	userData := userResponseFormatter(newUser, authToken)

	// TODO development: ApiResponseFormatter
	response := helper.ResponseFormatter(http.StatusOK, "success", "user registered", userData)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) PostUserLogin(c echo.Context) error {
	req := new(loginRequest)
	if err := c.Bind(req); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", "invalid request", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	user, err := h.userServices.signIn(req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusNotFound, "fail", err.Error(), nil)
		return c.JSON(http.StatusNotFound, response)
	}

	// TODO auth-token
	authToken := "123"
	userData := userResponseFormatter(user, authToken)

	response := helper.ResponseFormatter(http.StatusOK, "success", "user authenticated", userData)

	return c.JSON(http.StatusOK, response)
}

// TODO error-handling
func (h *userHandler) GetUsers(c echo.Context) error {
	if findByEmail := c.QueryParam("email"); findByEmail != "" {
		email := c.QueryParam("email")

		user, err := h.userServices.FetchUserByEmail(email)
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.M{"message": err.Error()})
		} else {
			response := user
			return c.JSON(http.StatusOK, response)
		}

	}
	// TODO response-formatter
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
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", "invalid request", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"fail": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	response, _ := h.userServices.UpdateUser(id, req)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.userServices.DeleteUser(id); err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", err, nil)
		return c.JSON(http.StatusOK, response)
	}
	message := fmt.Sprintf("user %d was deleted", id)
	response := helper.ResponseFormatter(http.StatusOK, "success", message, nil)
	return c.JSON(http.StatusOK, response)
}
