// Package utils implements generic utilities for api
package utils

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// HTTPResponse struct defines fields for http error
type HTTPResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// SendHTTPResponse sends http response to client
func SendHTTPResponse(data interface{}, status int, w http.ResponseWriter) {
	w.WriteHeader(status)

	newHTTPResponse := HTTPResponse{
		Data:   data,
		Status: status,
	}

	log.Info(newHTTPResponse)
	json.NewEncoder(w).Encode(newHTTPResponse)
}
