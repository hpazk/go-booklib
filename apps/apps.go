package apps

import (
	"github.com/hpazk/go-booklib/apps/user"
	"github.com/jinzhu/gorm"
)

func AppInit(db *gorm.DB) {
	userapp := user.UserInit(db)
	userapp.InitApp()
}
