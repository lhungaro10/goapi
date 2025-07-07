package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditOpenningHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"PUT opening",
	})
}