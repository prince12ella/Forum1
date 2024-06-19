package controllers

import (
	"forum/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
	"strconv"
)

func ShowPosts(c *gin.Context) {
	var posts []models.Post
	db.Preload("User").Preload("Category").Find(&posts)
	c.HTML(http.StatusOK, "index.html", gin.H{"posts": posts})
}

func ShowNewPostPage(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	var categories []models.Category
	db.Find(&categories)
	c.HTML(http.StatusOK, "post.html", gin.H{"categories": categories})
}

func CreatePost(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}

	title := c.PostForm("title")
	content := c.PostForm("content")
	categoryID, _ := strconv.Atoi(c.PostForm("category"))

	post := models.Post{Title: title, Content: content, UserID: userID.(uint), CategoryID: uint(categoryID)}
	db.Create(&post)
	c.Redirect(http.StatusMovedPermanently, "/")
}

func ShowPost(c *gin.Context) {
	postID, _ := strconv.Atoi(c.Param("id"))

	var post models.Post
	db.Preload("User").Preload("Comments.User").First(&post, postID)

	c.HTML(http.StatusOK, "post.html", gin.H{"post": post})
}

func CreateComment(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}

	postID, _ := strconv.Atoi(c.Param("id"))
	content := c.PostForm("content")

	comment := models.Comment{Content: content, PostID: uint(postID), UserID: userID.(uint)}
	db.Create(&comment)
	c.Redirect(http.StatusMovedPermanently, "/post/"+strconv.Itoa(postID))
}
func ShowPostsByCategory(c *gin.Context) {
	categoryID, _ := strconv.Atoi(c.Param("category_id"))

	var posts []models.Post
	db.Preload("User").Preload("Category").Where("category_id = ?", categoryID).Find(&posts)

	c.HTML(http.StatusOK, "index.html", gin.H{"posts": posts})
}
