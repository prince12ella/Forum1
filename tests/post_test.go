package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"forum/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	// Configurer l'environnement de test
	r := gin.Default()
	// Ajouter les routes et handlers nécessaires

	// Créer une requête HTTP de test
	req, _ := http.NewRequest("POST", "/post/new", strings.NewReader("title=Test Post&content=This is a test post&category_id=1"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Enregistrer la réponse
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Vérifier le statut de la réponse
	assert.Equal(t, http.StatusMovedPermanently, w.Code)

	// Vérifier si le post a été ajouté dans la base de données
	var post models.Post
	db.Where("title = ?", "Test Post").First(&post)
	assert.NotNil(t, post)
	assert.Equal(t, "This is a test post", post.Content)
}
