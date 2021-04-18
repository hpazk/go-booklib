package apps

import (
	"fmt"

	"github.com/hpazk/go-booklib/apps/book"
	"github.com/hpazk/go-booklib/apps/user"
	"github.com/hpazk/go-booklib/database"
)

func AppInit() {
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
	user := user.UserInit(db)
	book := book.BookInit(db)
	user.InitApp()
	book.InitApp()
}
