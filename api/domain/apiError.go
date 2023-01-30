package domain

import (
	"fmt"
	"net/http"
)

type ApiError interface {
	Message() string
	Code() string
	Status() int
	Cause() string
	Error() string
}

type apiError struct {
	ErrorMessage string `json:"message"`
	ErrorCode    string `json:"error"`
	ErrorStatus  int    `json:"status"`
	ErrorCause   string `json:"cause"`
}

func (e apiError) Code() string {
	return e.ErrorCode
}

func (e apiError) Error() string {
	return fmt.Sprintf("Message: %s;Error Code: %s;Status: %d;Cause: %v", e.ErrorMessage, e.ErrorCode, e.ErrorStatus, e.ErrorCause)
}

func (e apiError) Status() int {
	return e.ErrorStatus
}

func (e apiError) Cause() string {
	return e.ErrorCause
}

func (e apiError) Message() string {
	return e.ErrorMessage
}

func NewApiError(message string, error string, status int, cause string) ApiError {
	return apiError{message, error, status, cause}
}

func NewBadRequestApiError(message string) ApiError {
	return apiError{message, "bad_request", http.StatusBadRequest, ""}
}

func NewInternalServerApiError(message string, err error) ApiError {
	return apiError{message, "internal_server_error", http.StatusInternalServerError, ""}
}
