package user

import (
	"github.com/hpazk/go-booklib/auth"
	"github.com/hpazk/go-booklib/helper"
	"github.com/labstack/echo/v4"
)

type UserRoutes struct {
}

func (r *UserRoutes) Route() []helper.Route {
	// db := database.GetDbInstance()
	// db.AutoMigrate(User{})
	// userRepo := NewRepository(db)
	// userService := NewServices(userRepo)
	// authService := auth.NewAuthService()

	repos := UserRepository(&UsersStorage{})
	userservice := UserService(repos)
	authService := auth.AuthService()

	handler := UserHandler(userservice, authService)

	// TODO jwt: OK
	// TODO token-validation
	// fmt.Println(authService.GetAccessToken(1, "admin"))
	// token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Miwicm9sZSI6Im1lbWJlciIsImV4cCI6MTYxODY4OTQwMywiaWF0IjoxNjE4Njc4NjAzfQ.FSRmtEvBspQisJen2hk0MvPtVKEpWUQpZBvwHoKe8ow")
	// if err != nil {
	// 	fmt.Println("Error")
	// }

	// if token.Valid {
	// 	fmt.Println("Token is Valid")
	// } else {
	// 	fmt.Println("Token is Invalid")
	// }

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
			Method:     echo.POST,
			Path:       "/user/upload/photo",
			Handler:    handler.PostUserPhoto,
			Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
		},
		{
			Method:     echo.GET,
			Path:       "/users",
			Handler:    handler.GetUsers,
			Middleware: []echo.MiddlewareFunc{auth.JwtMiddleWare()},
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
