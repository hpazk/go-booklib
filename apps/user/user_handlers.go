package user

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/hpazk/go-booklib/auth"
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
	authServices auth.AuthServices
}

func UserHandler(userServices UserServices, authServices auth.AuthServices) *userHandler {
	return &userHandler{userServices, authServices}
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

	// TODO services:check-email-exist

	newUser, _ := h.userServices.signUp(req)

	// TODO development: authToken
	authToken, err := h.authServices.GetAccessToken(newUser.ID, newUser.Role)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", "something went wrong", nil)

		return c.JSON(http.StatusInternalServerError, response)
	}

	// TODO development: userData
	userData := userResponseFormatter(newUser, authToken)

	response := helper.ResponseFormatter(http.StatusOK, "success", "user registered", userData)
	return c.JSON(http.StatusOK, response)
}

// TODO check exist-email
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

	signedInUser, err := h.userServices.signIn(req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusNotFound, "fail", err.Error(), nil)
		return c.JSON(http.StatusNotFound, response)
	}

	// TODO auth-token
	authToken, err := h.authServices.GetAccessToken(signedInUser.ID, signedInUser.Role)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", "something went wrong", nil)

		return c.JSON(http.StatusInternalServerError, response)
	}
	userData := userLoginResponseFormatter(signedInUser, authToken)

	response := helper.ResponseFormatter(http.StatusOK, "success", "user authenticated", userData)

	return c.JSON(http.StatusOK, response)
}

// TODO error-handling
func (h *userHandler) PostUserPhoto(c echo.Context) error {

	// TODO jwt: userId
	// TODO image-validation
	id := 1
	user, err := h.userServices.FetchUserById(id)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusNotFound, "fail", "user doesn't exist", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	// Source
	photo, err := c.FormFile("photo")
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", "file upload failed", helper.M{"is_uploaded": false})
		return c.JSON(http.StatusBadRequest, response)
	}

	src, err := photo.Open()
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", "file upload failed", helper.M{"is_uploaded": false})
		return c.JSON(http.StatusInternalServerError, response)
	}
	defer src.Close()

	ext := string(photo.Filename[len(photo.Filename)-3:])

	photoPath := fmt.Sprintf("public/images/%d-%s.%s", user.ID, user.Name, ext)

	// Destination
	dst, err := os.Create(photoPath)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", "file upload failed", helper.M{"is_uploaded": false})
		return c.JSON(http.StatusInternalServerError, response)
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", "file upload failed", helper.M{"is_uploaded": false})
		return c.JSON(http.StatusInternalServerError, response)
	}

	// Upload
	_, err = h.userServices.UploadPhoto(user, photoPath)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", "file upload failed", helper.M{"is_uploaded": false})
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.ResponseFormatter(http.StatusOK, "success", "photo succesfully uploaded", helper.M{"is_uploaded": true})

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
