// Package utils implements generic utilities for api
package utils

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// HTTPError struct defines fields for http error
type HTTPError struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

// SendHTTPError sends http error to client
func SendHTTPError(status int, err error, w http.ResponseWriter) {
	w.WriteHeader(status)
	
	newHTTPError := HTTPError{
		Status: status,
		Message: err.Error(),
	}
	
	log.Error(newHTTPError)
	json.NewEncoder(w).Encode(newHTTPError)
}