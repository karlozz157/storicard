package errors

import (
	"fmt"
	"net/http"
)

var ErrEmailInvalid = NewErrStori("email invalid", http.StatusInternalServerError)
var ErrEmailRequired = NewErrStori("email is required", http.StatusBadRequest)
var ErrInternal = NewErrStori("internal", http.StatusInternalServerError)

type ErrStori struct {
	Mesage     string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func NewErrStori(message string, statusCode int) *ErrStori {
	return &ErrStori{
		Mesage:     message,
		StatusCode: statusCode,
	}
}

func (e *ErrStori) Error() string {
	return fmt.Sprint(e.Mesage)
}

func ParseError(err error) (int, string) {
	message := "houston, we have a problem"
	statusCode := http.StatusInternalServerError

	if errStori, ok := err.(*ErrStori); ok {
		message = errStori.Mesage
		statusCode = errStori.StatusCode
	}

	return statusCode, message
}
