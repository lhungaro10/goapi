package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine){
	v1 := router.Group("/api/v1")
	{
		//server status
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		//Show Opportunity
		v1.GET("/opening", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{
				"message":"GET opening",
			})
		})
		//Create Opportunity
		v1.POST("/opening", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{
				"message":"POST opening",
			})
		})

		//Edit Opportunity
		v1.PUT("/opening", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{
				"message":"PUT opening",
			})
		})

		//Delete Opportunity
		v1.DELETE("/opening", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{
				"message":"DELETE opening",
			})
		})
		//Show All Opportunity
		v1.GET("/openings", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{
				"message":"GET openings",
			})
		})
	}
}
