package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"forum/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLikePost(t *testing.T) {
	// Configurer l'environnement de test
	r := gin.Default()
	// Ajouter les routes et handlers nécessaires

	// Créer un utilisateur et un post pour les tests
	user := models.User{Email: "test@example.com", Username: "testuser", Password: "password"}
	db.Create(&user)
	post := models.Post{Title: "Test Post", Content: "This is a test post", UserID: user.ID}
	db.Create(&post)

	// Simuler la session de l'utilisateur
	r.Use(func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("user_id", user.ID)
		session.Save()
	})

	// Créer une requête HTTP de test
	req, _ := http.NewRequest("POST", "/post/"+strconv.Itoa(int(post.ID))+"/like", nil)

	// Enregistrer la réponse
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Vérifier le statut de la réponse
	assert.Equal(t, http.StatusMovedPermanently, w.Code)

	// Vérifier si le like a été ajouté dans la base de données
	var like models.Like
	db.Where("post_id = ? AND user_id = ?", post.ID, user.ID).First(&like)
	assert.NotNil(t, like)
}
