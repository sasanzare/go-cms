package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sasanzare/go-cms/controllers"
	"github.com/sasanzare/go-cms/middleware"
)

func SetupPostRoutes(r *gin.Engine) {
	posts := r.Group("/api/posts")
	{
		posts.GET("", controllers.ListPosts)
		posts.GET("/:id", controllers.GetPost)
		posts.POST("", middleware.AuthMiddleware(), controllers.CreatePost)
		posts.PUT("/:id", middleware.AuthMiddleware(), controllers.UpdatePost)
	}
}
