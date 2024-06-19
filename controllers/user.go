package controllers

import (
    "net/http"
    "strconv"

    "forum/models"

    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
)

// ShowPostsByUser affiche les posts créés par l'utilisateur connecté
func ShowPostsByUser(c *gin.Context) {
    session := sessions.Default(c)
    userID := session.Get("user_id")
    if userID == nil {
        c.Redirect(http.StatusMovedPermanently, "/login")
        return
    }

    var posts []models.Post
    models.DB.Preload("User").Preload("Category").Where("user_id = ?", userID).Find(&posts)

    c.HTML(http.StatusOK, "index.html", gin.H{"posts": posts})
}

// ShowPostsLikedByUser affiche les posts aimés par l'utilisateur connecté
func ShowPostsLikedByUser(c *gin.Context) {
    session := sessions.Default(c)
    userID := session.Get("user_id")
    if userID == nil {
        c.Redirect(http.StatusMovedPermanently, "/login")
        return
    }

    var likes []models.Like
    models.DB.Preload("Post").Preload("Post.User").Preload("Post.Category").Where("user_id = ?", userID).Find(&likes)

    var posts []models.Post
    for _, like := range likes {
        posts = append(posts, like.Post)
    }

    c.HTML(http.StatusOK, "index.html", gin.H{"posts": posts})
}
