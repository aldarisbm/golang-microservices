package controllers

import (
	"net/http"
	"strconv"

	"github.com/aldarisbm/golang-microservices/services"
	"github.com/aldarisbm/golang-microservices/utils"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UsersService.GetUser(userID)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, user)
}
