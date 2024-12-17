package apiError

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrBadRequest          = errors.New("bad request")
	ErrNotFound            = errors.New("not found")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
	ErrTooManyReq          = errors.New("too many requests")
	ErrConflict            = errors.New("conflict")
	ErrInternalServerError = errors.New("internal server error")
	ErrServiceUnavaliable  = errors.New("service unavaliable")
)

type ApiErr interface {
	Status() int
	Message() string
	Error() string
}

type ApiError struct {
	ErrorStatus  int         `json:"status"`
	ErrorMessage string      `json:"message"`
	Errors       interface{} `json:"error"`
}

func (a ApiError) Status() int {
	return a.ErrorStatus
}

func (a ApiError) Message() string {
	return a.ErrorMessage
}

func (a ApiError) Error() string {
	return fmt.Sprintf("%v", a.Errors)
}

func NewApiError(status int, message string, error interface{}) ApiErr {
	return ApiError{
		ErrorStatus:  status,
		ErrorMessage: message,
		Errors:       error,
	}
}

func NewBadRequestError(error interface{}) ApiErr {
	return ApiError{
		ErrorStatus:  http.StatusBadRequest,
		ErrorMessage: ErrBadRequest.Error(),
		Errors:       error,
	}
}

func NewNotFoundError(error interface{}) ApiErr {
	return ApiError{
		ErrorStatus:  http.StatusNotFound,
		ErrorMessage: ErrNotFound.Error(),
		Errors:       error,
	}
}

func NewUnauthorizedError(error interface{}) ApiErr {
	return ApiError{
		ErrorStatus:  http.StatusUnauthorized,
		ErrorMessage: ErrUnauthorized.Error(),
		Errors:       error,
	}
}

func NewForbiddenError(error interface{}) ApiErr {
	return ApiError{
		ErrorStatus:  http.StatusForbidden,
		ErrorMessage: ErrForbidden.Error(),
		Errors:       error,
	}
}

func NewTooManyRequestsError(error interface{}) ApiErr {
	return ApiError{
		ErrorStatus:  http.StatusTooManyRequests,
		ErrorMessage: ErrTooManyReq.Error(),
		Errors:       error,
	}
}

func NewConflictError(error interface{}) ApiErr {
	return ApiError{
		ErrorStatus:  http.StatusConflict,
		ErrorMessage: ErrConflict.Error(),
		Errors:       error,
	}
}

func NewInternalServerError(error interface{}) ApiErr {
	return ApiError{
		ErrorStatus:  http.StatusInternalServerError,
		ErrorMessage: ErrInternalServerError.Error(),
		Errors:       error,
	}
}

func NewServiceUnavaliableError(error interface{}) ApiErr {
	return ApiError{
		ErrorStatus:  http.StatusServiceUnavailable,
		ErrorMessage: ErrServiceUnavaliable.Error(),
		Errors:       error,
	}
}
