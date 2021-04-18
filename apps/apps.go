package apps

import (
	"github.com/hpazk/go-booklib/apps/book"
	"github.com/hpazk/go-booklib/apps/user"
	"github.com/jinzhu/gorm"
)

func AppInit(db *gorm.DB) {
	user := user.UserInit(db)
	book := book.BookInit(db)
	user.InitApp()
	book.InitApp()
}
