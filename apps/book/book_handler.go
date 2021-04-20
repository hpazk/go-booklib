package book

import (
	"fmt"
	"net/http"
	"strconv"

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
	response := helper.ResponseFormatter(http.StatusOK, "success", "book stored", bookData)
	return c.JSON(http.StatusOK, response)
}

func GetBook(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "book-newest"})
}

func (h *bookHandler) GetBooks(c echo.Context) error {
	category := c.QueryParam("category")

	if category != "" {
		// categoryId, _ := h.bookServices.FetchByCategory(category)
		books, err := h.bookServices.FetchBooksByCategory(category)
		if err != nil {
			response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", err.Error(), nil)
			return c.JSON(http.StatusInternalServerError, response)
		}

		booksData := booksResponseFormatter(books)

		response := helper.ResponseFormatter(http.StatusOK, "success", "book successfully fetched", booksData)
		return c.JSON(http.StatusOK, response)
	}

	books, err := h.bookServices.FetchBooks()
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	booksData := booksResponseFormatter(books)

	response := helper.ResponseFormatter(http.StatusOK, "success", "book successfully fetched", booksData)

	return c.JSON(http.StatusOK, response)
}

func (h *bookHandler) GetNewestBooks(c echo.Context) error {
	books, err := h.bookServices.FetchNewestBooks()
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	booksData := booksResponseFormatter(books)

	response := helper.ResponseFormatter(http.StatusOK, "success", "newest book successfully fetched", booksData)

	return c.JSON(http.StatusOK, response)
}

func (h *bookHandler) PutBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := new(request)

	fmt.Println(id)

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

	editedBook, err := h.bookServices.EditBook(uint(id), req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	bookData := helper.ResponseFormatter(http.StatusOK, "success", "book successfully updated", editedBook)

	return c.JSON(http.StatusOK, bookData)
}

func (h *bookHandler) DeleteBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.bookServices.RemoveBook(uint(id))
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.ResponseFormatter(http.StatusOK, "success", "book was deleted", nil)

	return c.JSON(http.StatusOK, response)
}
