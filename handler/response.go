package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, message string){
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": message,
		"errorCode": code,
	})
}

func SendSuccess(ctx *gin.Context, operation string, data interface{}){
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation for handle: %s successful", operation),
		"data": data,
	})
}