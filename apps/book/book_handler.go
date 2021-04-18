package book

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/hpazk/go-booklib/helper"
	"github.com/labstack/echo/v4"
)

type bookHandler struct {
	bookServices BookServices
}

func BookHandler(userServices BookServices) *bookHandler {
	return &bookHandler{userServices}
}

func (h *bookHandler) PostBook(c echo.Context) error {
	accessToken := c.Get("user").(*jwt.Token)
	claims := accessToken.Claims.(jwt.MapClaims)
	role := claims["user_role"]

	if role != "admin" {
		response := helper.ResponseFormatter(http.StatusUnauthorized, "fail", "Please provide valid credentials", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}
	req := new(request)

	// Check request
	if err := c.Bind(req); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", "invalid request", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	// Validate request
	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	// SignUp service
	newBook, err := h.bookServices.AddBook(req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	// Format data
	bookData := bookResponseFormatter(newBook)

	// Passed response
	response := helper.ResponseFormatter(http.StatusOK, "success", "user registered", bookData)
	return c.JSON(http.StatusOK, response)
}

func UpdateBook(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "book-newest"})
}

func GetBook(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "book-newest"})
}

func GetBooks(c echo.Context) error {
	// TODO 1: get books
	// TODO 1: get books by category
	category := c.QueryParam("category")
	fmt.Println(category)

	return c.JSON(http.StatusOK, helper.M{"message": "book-catalog"})
}

func GetNewestBooks(c echo.Context) error {
	// TODO 2: get books by newest year

	return c.JSON(http.StatusOK, helper.M{"message": "book-newest"})
}

func DeleteBook(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "book-delete"})
}
