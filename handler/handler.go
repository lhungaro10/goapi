package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpenningHandler(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
		"message":"GET opening",
	})
}

func CreateOpenningHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"POST opening",
	})
}

func EditOpenningHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"PUT opening",
	})
}

func DeleteOpenningHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"DELETE opening",
	})
}
func ShowAllOpenningsHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"GET openings",
	})
}