package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/lhungaro10/goapi/docs"
	"github.com/lhungaro10/goapi/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine){
	//initialize handler
	handler.InitializeHandler()
	const BASE_PATH string = "/api/v1"
	docs.SwaggerInfo.BasePath = BASE_PATH;

	v1 := router.Group(BASE_PATH)
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
