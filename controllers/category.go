// controllers/category.go

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prince12ella/Forum1/db"
	"github.com/prince12ella/Forum1/models"
)

func GetCategories(c *gin.Context) {
	// Initialisation de la connexion à la base de données
	database, err := db.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not available"})
		return
	}
	defer database.Close()

	var categories []models.Category
	if err := database.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}
