package controllers

import (
    "net/http"
    "strconv"

    "github.com/prince12ella/Forum1/models"
db "github.com/prince12ella/Forum1/db"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
)

// LikePost permet à un utilisateur de liker un post
func LikePost(c *gin.Context) {
    session := sessions.Default(c)
    userID := session.Get("user_id")
    if userID == nil {
        c.Redirect(http.StatusMovedPermanently, "/login")
        return
    }

    postID, _ := strconv.Atoi(c.Param("id"))

    like := models.Like{PostID: uint(postID), UserID: userID.(uint)}
    db.Create(&like)
    c.Redirect(http.StatusMovedPermanently, "/post/"+strconv.Itoa(postID))
}

// DislikePost permet à un utilisateur de disliker un post
func DislikePost(c *gin.Context) {
    session := sessions.Default(c)
    userID := session.Get("user_id")
    if userID == nil {
        c.Redirect(http.StatusMovedPermanently, "/login")
        return
    }

    postID, _ := strconv.Atoi(c.Param("id"))

    dislike := models.Dislike{PostID: uint(postID), UserID: userID.(uint)}
    models.DB.Create(&dislike)
    c.Redirect(http.StatusMovedPermanently, "/post/"+strconv.Itoa(postID))
}
