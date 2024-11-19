package user

import (
	"url-shortening-api/internal/link"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string      `gorm:"column:uuid"`
	Username string      `gorm:"column:username"`
	Email    string      `gorm:"column:email"`
	Password string      `gorm:"column:password"`
	Links    []link.Link `gorm:"foreignkey:user_id;references:uuid"`
}

func (u *User) TableName() string {
	return "users"
}
