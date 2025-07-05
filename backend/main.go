package main

import (
	"github.com/sasanzare/go-cms/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupPostRoutes(r)
	r.Run(":8000")
}
