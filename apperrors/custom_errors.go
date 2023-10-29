package apperrors

import (
	"fmt"
	"net/http"
)

var (
	ErrUserNotFound  = NewHTTPError(http.StatusNotFound, "User not found")
	ErrInvalidRequest = NewHTTPError(http.StatusBadRequest, "Invalid request")
	// Добавьте здесь другие пользовательские ошибки
)

func NewHTTPError(status int, message string) HTTPError {
	return HTTPError{
		Status:  status,
		Message: message,
	}
}

type ErrorResponse struct {
    Message string `json:"message"`
}

type HTTPError struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
}

func (e HTTPError) Error() string {
	return e.Message
}

func (e HTTPError) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.Status)
	return nil
}

func NewInvalidRequestError(err error) HTTPError {
	return HTTPError{
		Status:  http.StatusBadRequest,
		Message: fmt.Sprintf("Invalid request: %v", err),
	}
}
