package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lhungaro10/goapi/schemas"
)

//@BasePath /api/v1
//@Summary Create Opening
//@Description Create a new Job Opening
//@Tags Opening
//@Accept json
//@Produce json
//@Param request body CreateOpeningRequest true "request body"
//@Success 200 {object} CreateOpeneningReponse
//@Failure 400 {object} ErrorResponse
//@Failure 500 {object} ErrorResponse
//@Router /opening [post]
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