package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lhungaro10/goapi/schemas"
)

func EditOpenningHandler(c *gin.Context){
	id := c.Query("id")
	request := UpdateOpeningRequest{}
	c.BindJSON(&request)
		
	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation error: %s", err.Error())
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if id == "" {
		SendError(c, 400, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		SendError(c, http.StatusNotFound, fmt.Sprintf("Opening for id %s not found", id))
		return
	}
	
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err!= nil {
		SendError(c, http.StatusInternalServerError, fmt.Sprintf("Error updating opening: %s", err))
		return
	}

	SendSuccess(c, "edit-opening", opening)
}