package apperror

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrNotFound is used when server needs to response with 404 Not Found status code.
	ErrNotFound = NewAppError(
		http.StatusNotFound,
		"requested resource is not found",
		"maybe you have an error in your request or requested resource does not exist",
	)

	// ErrNoRows is used when no rows returned from storage.
	ErrNoRows = errors.New("no rows")

	// ErrValidationFailed is used when input validation failed.
	ErrValidationFailed = errors.New("input validation failed. you've provided invalid values")

	// ErrEmailTaken is used when the user is being created and given email is already taken.
	ErrEmailTaken = errors.New("email already taken")

	// ErrWrongPassword is used when client provided wrong password.
	ErrWrongPassword = errors.New("wrong email or password")

	// ErrInvalidRequestBody is used when client sends invalid request body.
	ErrInvalidRequestBody = errors.New("invalid request body")
)

// AppError describes a structure of an error response in JSON format.
type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developerMessage,omitempty"`
	HttpCode         int    `json:"code,omitempty"`
} // @name ErrorResponse

// NewAppError returns a new AppError instance.
func NewAppError(code int, message, developerMessage string) *AppError {
	return &AppError{
		Err:              fmt.Errorf(message),
		Message:          message,
		DeveloperMessage: developerMessage,
		HttpCode:         code,
	}
}

// Error returns a string representation of an error.
func (ae *AppError) Error() string {
	return ae.Err.Error()
}

// BadRequestError returns a new AppError instance
// with 400 Bad Request status code.
func BadRequestError(message, developerMessage string) *AppError {
	return NewAppError(http.StatusBadRequest, message, developerMessage)
}

// InternalError returns a new AppError instance
// with 500 Internal Server Error status code.
func InternalError(message, developerMessage string) *AppError {
	return NewAppError(http.StatusInternalServerError, message, developerMessage)
}
