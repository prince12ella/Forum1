package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/prince12ella/Forum1/db"
	"github.com/prince12ella/Forum1/models"
)

func GetCategories(c *gin.Context) {
    db, err := db.InitDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not available"})
        return
    }
    defer db.Close()

    var categories []models.Category
    if err := db.Find(&categories).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, categories)
}
