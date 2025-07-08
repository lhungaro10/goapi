package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lhungaro10/goapi/schemas"
)

func DeleteOpenningHandler(c *gin.Context){
	id := c.Query("id")

	if id == "" {
		SendError(c, 400, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	//find opening by id
	if err := db.First(&opening, id).Error; err != nil {
		SendError(c, http.StatusNotFound, fmt.Sprintf("opening for id %s not found", id))
		return
	}

	//delete opening
	if err := db.Delete(&opening).Error; err != nil {
		SendError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting opening: %s", err))
		return
	}

	SendSuccess(c, "delete-opening", opening)
}