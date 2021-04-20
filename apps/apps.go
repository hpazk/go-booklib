package apps

import (
	"fmt"

	"github.com/hpazk/go-booklib/apps/book"
	"github.com/hpazk/go-booklib/apps/user"
	"github.com/hpazk/go-booklib/database"
	"github.com/hpazk/go-booklib/helper"
	"github.com/hpazk/go-booklib/routes"
	"github.com/labstack/echo/v4"
)

func AppInit(e *echo.Echo) {
	// Database
	db := database.GetDbInstance()
	dbMigration := database.GetMigrations(db)
	err := dbMigration.Migrate()
	if err == nil {
		fmt.Println("Migrations did run successfully")
	} else {
		fmt.Println("migrations failed.", err)
	}

	// Apps
	userApp := user.InitApp(db)
	bookApp := book.InitApp(db)
	userApp.UseApp()
	bookApp.UseApp()

	// Route App Handler
	handlers := []helper.Handler{
		&user.UserApp{},
		&book.BookApp{},
	}

	routes.DefineApiRoutes(e, handlers)
}
