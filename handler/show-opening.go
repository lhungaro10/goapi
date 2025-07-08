package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lhungaro10/goapi/schemas"
)

func ShowOpenningHandler(c *gin.Context){
	id := c.Query("id")

	if id == "" {
		SendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		SendError(c, http.StatusNotFound, fmt.Sprintf("opening for id %s not found", id))
		return
	}

	SendSuccess(c, "show-opening", opening)
}