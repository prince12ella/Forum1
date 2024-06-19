package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prince12ella/Forum1/controllers"
	db "github.com/prince12ella/Forum1/db"
)

func main() {
	r := gin.Default()

	// Initialise la base de données
	_, err := db.InitDB()
	if err != nil {
		panic("failed to connect database")
	}
	defer db.DB.Close()

	// Injecte la base de données dans le contexte Gin
	r.Use(func(c *gin.Context) {
		c.Set("db", db.DB)
		c.Next()
	})

	// Définir les routes
	r.GET("/categories", controllers.GetCategories)
	r.POST("/like/:id", controllers.LikePost)
	r.POST("/post/:id", controllers.CreatePost) // Exemple, ajustez selon vos besoins

	r.Run(":8080") // Démarre le serveur sur le port 8080
}
