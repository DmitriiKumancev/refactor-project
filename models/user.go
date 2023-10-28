package models

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
    Err            error  `json:"-"`
    HTTPStatusCode int    `json:"-"`

    StatusText     string `json:"status"`
    AppCode        int64  `json:"code,omitempty"`
    ErrorText      string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
    render.Status(r, e.HTTPStatusCode)
    return nil
}

func ErrInvalidRequest(err error) render.Renderer {
    return &ErrResponse{
        Err:            err,
        HTTPStatusCode: 400,
        StatusText:     "Invalid request.",
        ErrorText:      err.Error(),
    }
}

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
    Email       string `json:"email"`
}

func (u *UpdateUserRequest) Bind(r *http.Request) error {
    return nil
}