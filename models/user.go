package models

import (
	"errors"
	"net/http"
)

type User struct {
    CreatedAt   string `json:"created_at"`
    DisplayName string `json:"display_name"`
    Email       string `json:"email"`
}

type UserList map[string]User

type UserStore struct {
    Increment int      `json:"increment"`
    List      UserList `json:"list"`
}

type CreateUserRequest struct {
    DisplayName string `json:"display_name"`
    Email       string `json:"email"`
}

func (c *CreateUserRequest) Bind(r *http.Request) error {
    if c.DisplayName == "" {
        return errors.New("display_name is required")
    }

    if c.Email == "" {
        return errors.New("email is required")
    }

    return nil 
}

type UpdateUserRequest struct {
    DisplayName string `json:"display_name"`
}

func (u *UpdateUserRequest) Bind(r *http.Request) error {
    return nil
}