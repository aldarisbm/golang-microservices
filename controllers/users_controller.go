package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aldarisbm/golang-microservices/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIDParam := req.URL.Query().Get("user_id")
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		// just return the Bad Request to the client
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("user_id must be a number"))
		return
	}

	user, err := services.GetUser(userID)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		//Handle the err
		return
	}
	//return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
