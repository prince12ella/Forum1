package models

// User represents a user model.
type User struct {
    ID       uint   `gorm:"primary_key"`
    Email    string `gorm:"unique;not null"`
    Username string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
}
