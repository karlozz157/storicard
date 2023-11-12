package errors

import (
	"fmt"
	"net/http"
)

var ErrInternal = NewErrStori("internal", http.StatusInternalServerError)

type ErrStori struct {
	Mesage     string
	StatusCode int
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
