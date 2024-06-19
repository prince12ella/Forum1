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

func TestUserRegistration(t *testing.T) {
	// Configurer l'environnement de test
	r := gin.Default()
	// Ajouter les routes et handlers nécessaires

	// Créer une requête HTTP de test
	req, _ := http.NewRequest("POST", "/register", strings.NewReader("email=test@example.com&username=testuser&password=password"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Enregistrer la réponse
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Vérifier le statut de la réponse
	assert.Equal(t, http.StatusMovedPermanently, w.Code)

	// Vérifier si l'utilisateur a été ajouté dans la base de données
	var user models.User
	db.Where("email = ?", "test@example.com").First(&user)
	assert.NotNil(t, user)
	assert.Equal(t, "testuser", user.Username)
}
