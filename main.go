package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prince12ella/Forum1/controllers"
	"github.com/prince12ella/Forum1/db"
)

func main() {
	r := gin.Default()

	// Initialise la base de données
	dbInstance, err := db.InitDB()
	if err != nil {
		panic("failed to connect database")
	}
	defer dbInstance.Close()

	// Injecte la base de données dans le contexte
	r.Use(func(c *gin.Context) {
		c.Set("db", dbInstance)
	})

	// Définir les routes
	r.GET("/categories", controllers.GetCategories)

	r.Run(":8080") // Démarre le serveur sur le port 8080
}
