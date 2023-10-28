package api

import (
	"errors"
	"net/http"
)

var (
	ErrUserNotFound = errors.New("user_not_found")
)

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	//
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	//
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	//
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//
}
