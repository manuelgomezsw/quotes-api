package apierror

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiError interface {
	Message() string
	Code() string
	Status() int
	Cause() []Cause
	Error() string
}

type apiErr struct {
	ErrorMessage string  `json:"message"`
	ErrorCode    string  `json:"error"`
	ErrorStatus  int     `json:"status"`
	ErrorCause   []Cause `json:"cause"`
}

type Cause struct {
	Domain     string   `json:"domain"`
	CauseID    int64    `json:"cause_id"`
	Type       string   `json:"type"`
	Code       string   `json:"code"`
	References []string `json:"references"`
	Message    string   `json:"message"`
}

func (e apiErr) Code() string {
	return e.ErrorCode
}

func (e apiErr) Error() string {
	return fmt.Sprintf("Message: %s;Error Code: %s;Status: %d;Cause: %v", e.ErrorMessage, e.ErrorCode, e.ErrorStatus, e.ErrorCause)
}

func (e apiErr) Status() int {
	return e.ErrorStatus
}

func (e apiErr) Cause() []Cause {
	return e.ErrorCause
}

func (e apiErr) Message() string {
	return e.ErrorMessage
}

func NewApiError(message string, error string, status int, cause []Cause) ApiError {
	return apiErr{message, error, status, cause}
}

func NewNotFoundApiError(message string, causeList []Cause) ApiError {
	return apiErr{message, "not_found", http.StatusNotFound, causeList}
}

func NewBadRequestApiError(message string, causeList []Cause) ApiError {
	return apiErr{message, "bad_request", http.StatusBadRequest, causeList}
}

func NewValidationApiError(message string, causeList []Cause) ApiError {
	return apiErr{message, "validation_error", http.StatusBadRequest, causeList}
}

func NewInternalServerApiError(message string, err error) ApiError {
	var cause []Cause
	if err != nil {
		cause = append(cause, Cause{Message: err.Error()})
	}
	return apiErr{message, "internal_server_error", http.StatusInternalServerError, cause}
}

func NewForbiddenApiError(message string) ApiError {
	return apiErr{message, "forbidden", http.StatusForbidden, []Cause{}}
}

func NewUnauthorizedApiError(message string) ApiError {
	return apiErr{message, "unauthorized_scopes", http.StatusUnauthorized, []Cause{}}
}

func NewApiErrorFromBytes(data []byte) (ApiError, error) {
	err := apiErr{}
	e := json.Unmarshal(data, &err)
	return err, e
}

func NewConflictApiError(message string) ApiError {
	return apiErr{message, "duplicated_key", http.StatusConflict, []Cause{}}
}
