package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	PostID  uint   `gorm:"not null"`
	Post    Post
	UserID  uint   `gorm:"not null"`
	User    User
}
