package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lhungaro10/goapi/handler"
)

func initializeRoutes(router *gin.Engine){
	//initialize handler
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		//server status
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		//Show Opportunity
		v1.GET("/opening", handler.ShowOpenningHandler)
		//Create Opportunity
		v1.POST("/opening", handler.CreateOpenningHandler)

		//Edit Opportunity
		v1.PUT("/opening", handler.EditOpenningHandler)

		//Delete Opportunity
		v1.DELETE("/opening", handler.DeleteOpenningHandler)
		//Show All Opportunity
		v1.GET("/openings", handler.ShowAllOpenningsHandler)
	}
}
