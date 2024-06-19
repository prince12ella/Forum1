package models

import (
	"time"
)

// Post représente un post dans le forum
type Post struct {
	ID        uint      `gorm:"primary_key"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"type:text;not null"`
	UserID    uint      `gorm:"not null"`
	CategoryID uint     `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User      User
	Category  Category
	Comments  []Comment
	Likes     []Like
	Dislikes  []Dislike
}
