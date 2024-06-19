package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *gorm.DB
)

// InitDB initialise la connexion à la base de données SQLite
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "db")
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}
