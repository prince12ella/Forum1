package models

// Like représente un like pour un post
type Like struct {
	ID        uint      `gorm:"primary_key"`
	PostID    uint      `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Post Post
	User User
}
