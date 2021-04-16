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

	// userHandler := NewHandler(userService, authService)

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/registration",
			Handler: PostUserRegistration,
		},
		{
			Method:  echo.GET,
			Path:    "/users",
			Handler: GetUsers,
		},
		{
			Method:  echo.GET,
			Path:    "/users/:id",
			Handler: GetUser,
		},
		{
			Method:  echo.PUT,
			Path:    "/users/:id",
			Handler: PutUser,
		},
		{
			Method:  echo.DELETE,
			Path:    "/users/:id",
			Handler: DeleteUser,
		},
	}
}
