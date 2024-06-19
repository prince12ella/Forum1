package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

func InitDB() (*gorm.DB, error) {
    db, err := gorm.Open("sqlite3", "forum.db")
    if err != nil {
        return nil, err
    }

    return db, nil
}
