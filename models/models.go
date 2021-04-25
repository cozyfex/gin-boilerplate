package models

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	var err error
	err = db.AutoMigrate(&Product{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		return
	}
}

type Product struct {
	gorm.Model
	Code  string
	Price int
}

// region User

type User struct {
	gorm.Model
	Username string
	Password string
	Nickname string
}

func (user *User) getOutput() string {
	return user.Nickname + "(" + user.Username + ")"
}

// endregion
