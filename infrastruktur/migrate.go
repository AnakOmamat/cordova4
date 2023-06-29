package infrastruktur

import (
	"cordova4/student"
	"cordova4/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	db.Debug().AutoMigrate(
		&user.User{},
		&student.Student{},
	)
}
