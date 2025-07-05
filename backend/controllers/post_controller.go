package controllers

import (
	"github.com/gin-gonic/gin"
)

func ListPosts(c *gin.Context) {
	c.JSON(200, gin.H{"message": "List posts"})
}

func GetPost(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get post"})
}

func CreatePost(c *gin.Context) {
	c.JSON(201, gin.H{"message": "Create post"})
}

func UpdatePost(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Update post"})
}

func ListPendingPosts(c *gin.Context) {
	c.JSON(200, gin.H{"message": "List pending posts"})
}

func ApprovePost(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Approve post"})
}

func RejectPost(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Reject post"})
}