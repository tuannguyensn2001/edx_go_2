package app

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type Error struct {
	StatusCode int         `json:"statusCode"`
	RootErr    error       `json:"-"`
	Message    string      `json:"message"`
	Log        string      `json:"log"`
	Key        string      `json:"key"`
	Data       interface{} `json:"data"`
}

func NewErrorResponse(message string, statusCode int, data interface{}, err error) *Error {

	logrus.WithFields(logrus.Fields{"error": err.Error()}).Warn(message)

	return &Error{
		Message:    message,
		StatusCode: statusCode,
		Data:       data,
	}
}

func (e *Error) RootError() error {
	if err, ok := e.RootErr.(*Error); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *Error) Error() string {
	return e.RootError().Error()
}

func ErrInvalidRequest(err error) *Error {
	return NewErrorResponse("Invalid Request", http.StatusBadRequest, nil, err)
}

func ErrInternalServer(err error) *Error {
	return NewErrorResponse("Error Server", http.StatusInternalServerError, nil, err)
}

func ErrEntityExisted(err error, message string) *Error {
	return NewErrorResponse(message, http.StatusConflict, nil, err)
}
