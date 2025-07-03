package main

import (
  "github.com/gin-gonic/gin"
//   "gorm.io/gorm"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{"message": "Hi"})
  })
  r.Run() // localhost:8080
}