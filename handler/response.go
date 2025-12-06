package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lhungaro10/goapi/schemas"
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

type ErrorResponse struct {
	Message string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeneningReponse struct {
	Message string `json:"message"`
	Data schemas.OpeningResponse `json:"data"`
}

type UpdateOpeneningReponse struct {
	Message string `json:"message"`
	Data schemas.OpeningResponse `json:"data"`
}

type DeleteOpeneningReponse struct {
	Message string `json:"message"`
	Data schemas.OpeningResponse `json:"data"`
}

type ShowOpeneningReponse struct {
	Message string `json:"message"`
	Data schemas.OpeningResponse `json:"data"`
}