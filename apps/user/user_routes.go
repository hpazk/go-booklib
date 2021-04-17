package user

import (
	"github.com/hpazk/go-booklib/helper"
	"github.com/labstack/echo/v4"
)

type UserRoutes struct{}

func (r *UserRoutes) Route() []helper.Route {
	// db := database.GetDbInstance()
	// db.AutoMigrate(User{})
	// userRepo := NewRepository(db)
	// userService := NewServices(userRepo)
	// authService := auth.NewAuthService()

	repos := UserRepository(&UsersStorage{})
	service := UserService(repos)
	handler := UserHandler(service)

	// userHandler := NewHandler(userService, authService)

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/registration",
			Handler: handler.PostUserRegistration,
		},
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: handler.PostUserLogin,
		},
		{
			Method:  echo.GET,
			Path:    "/users",
			Handler: handler.GetUsers,
		},
		{
			Method:  echo.GET,
			Path:    "/users/:id",
			Handler: handler.GetUser,
		},
		{
			Method:  echo.PUT,
			Path:    "/users/:id",
			Handler: handler.PutUser,
		},
		{
			Method:  echo.DELETE,
			Path:    "/users/:id",
			Handler: handler.DeleteUser,
		},
	}
}
