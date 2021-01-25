// Package handlers implements http handlers for api endpoints
package handlers

import (
	"net/http"

	"github.com/Akshit8/go-boilerplate/api/utils"
)

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, req *http.Request) {
	utils.SendHTTPResponse("from create user", 200, w)
}

// GetUser return a specific user
func GetUser(w http.ResponseWriter, req *http.Request) {
	utils.SendHTTPResponse("from get user", 200, w)
}

// Get returns all users
func Get(w http.ResponseWriter, req *http.Request) {
	utils.SendHTTPResponse("from get all users", 200, w)
}

// UpdateUser updates a specific user
func UpdateUser(w http.ResponseWriter, req *http.Request) {
	utils.SendHTTPResponse("from update user", 200, w)
}

// DeleteUser deletes a specific user
func DeleteUser(w http.ResponseWriter, req *http.Request) {
	utils.SendHTTPResponse("from delete user", 200, w)
}

