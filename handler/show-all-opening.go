package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lhungaro10/goapi/schemas"
)

func ShowAllOpenningsHandler(c *gin.Context){
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		SendError(c, http.StatusInternalServerError, "error fetching openings")
		return
	}

	SendSuccess(c, "show-all-openings", openings)
}