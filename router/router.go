package router

import "github.com/gin-gonic/gin"

func Initialize(){
	// Initialize Gin Router
	router := gin.Default()

	//Initialize Routes
	initializeRoutes(router)
	router.Run(":8080")
}