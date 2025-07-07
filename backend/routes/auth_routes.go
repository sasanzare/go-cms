package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	// Setup all main routes
	SetupPostRoutes(r)
	// Add other route setups here as needed
}
