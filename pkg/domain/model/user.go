package model

import "gorm.io/gorm"

type User struct {
	UserName string `gorm:"column:user_name;"`
	Password string
	gorm.Model
}

func (u *User) Table() string {
	return "users"
}
