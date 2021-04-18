package book

import (
	"github.com/hpazk/go-booklib/auth"
	"github.com/hpazk/go-booklib/helper"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type BookApp struct {
	Db *gorm.DB
}

func BookInit(db *gorm.DB) *BookApp {
	return &BookApp{db}
}

var handler *bookHandler

func (r *BookApp) InitApp() {
	repository := BookRepository(r.Db)
	bookservice := BookService(repository)

	handler = BookHandler(bookservice)
}

func (r *BookApp) Route() []helper.Route {
	return []helper.Route{
		{
			Method:     echo.POST,
			Path:       "/books",
			Handler:    handler.PostBook,
			Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		},
		{
			Method:  echo.GET,
			Path:    "/books",
			Handler: handler.GetBooks,
		},
		{
			Method:  echo.GET,
			Path:    "/books/newest",
			Handler: handler.GetNewestBooks,
		},
	}
}
