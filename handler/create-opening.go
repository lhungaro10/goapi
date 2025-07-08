package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lhungaro10/goapi/schemas"
)


func CreateOpenningHandler(c *gin.Context){
	request := CreateOpeningRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Error Validating request: %v", err)
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote: *request.Remote,
		Link: request.Link,
		Salary: request.Salary,
	};

	// create opening in database
	if err := db.Create(&opening).Error; err != nil {
		logger.ErrorF("Error creating opening: %v", err.Error())
		SendError(c, http.StatusInternalServerError, "Error creating opening")
		return
	}
	SendSuccess(c, "create-opening", opening)

}