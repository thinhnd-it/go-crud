package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	res := initializers.DB.Create(&post)

	if res.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"data": post,
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"data": posts,
	})
}

func PostShow(c *gin.Context) {
	var post models.Post
	postId := c.Param("postId")

	initializers.DB.Find(&post, postId)

	c.JSON(200, gin.H{
		"data": post,
	})
}

func PostUpdate(c *gin.Context) {
	var post models.Post
	var postReq models.Post
	c.Bind(&postReq)
	postId := c.Param("postId")

	initializers.DB.Find(&post, postId)
	initializers.DB.Model(&post).Updates(models.Post{
		Title: postReq.Title,
		Body:  postReq.Body,
	})

	c.JSON(200, gin.H{
		"data": post,
	})
}

func PostDelete(c *gin.Context) {
	postId := c.Param("postId")

	initializers.DB.Delete(&models.Post{}, postId)

	c.JSON(200, gin.H{
		"data": postId,
	})
}
