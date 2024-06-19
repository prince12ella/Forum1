package models

import (
	"time"
)

// Dislike représente un dislike pour un post
type Dislike struct {
	ID        uint      `gorm:"primary_key"`
	PostID    uint      `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Post Post
	User User
}
