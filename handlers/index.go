package handlers

import (
	"encoding/json"
	"net/http"
)

// IndexHandler is a response function for route mechanism
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	responsejson := successjson{}
	responsejson.Status = "welcome to notema"
	resp, _ := json.Marshal(responsejson)
	w.Write(resp)
}
