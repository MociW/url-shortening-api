package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string `gorm:"column:uuid"`
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func (u *User) TableName() string {
	return "users"
}
