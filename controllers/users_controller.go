package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aldarisbm/golang-microservices/services"
	"github.com/aldarisbm/golang-microservices/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIDParam := req.URL.Query().Get("user_id")
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)

		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userID)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)

		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
