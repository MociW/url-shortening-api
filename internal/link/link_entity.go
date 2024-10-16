package link

import (
	"time"
)

type Link struct {
	ID        uint      `gorm:"column:id"`
	UserId    string    `gorm:"column:user_id"`
	Link      string    `gorm:"column:link"`
	ShortLink string    `gorm:"column:short_link"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (l *Link) TableName() string {
	return "links"
}
