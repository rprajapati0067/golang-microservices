package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	AStatus  int    `json:"status"`
	AMessage string `json:"message"`
	AError   string `json:"error,omitempty"`
}

func (e *apiError) Status() int {
	return e.AStatus
}
func (e *apiError) Message() string {
	return e.AMessage
}
func (e *apiError) Error() string {
	return e.AError
}

func NewApiError(statusCode int, message string) ApiError {
	return &apiError{
		AStatus:  statusCode,
		AMessage: message,
	}
}

func NewApiErrFromBytes(body []byte) (ApiError, error) {
	var result apiError
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("invalid json body")
	}
	return &result, nil
}

func NewNotFoundApiError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusNotFound,
		AMessage: message,
	}
}
func NewInternalServerError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusInternalServerError,
		AMessage: message,
	}
}
func NewBadRequestError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusBadRequest,
		AMessage: message,
	}
}
