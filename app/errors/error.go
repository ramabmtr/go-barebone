package errors

import (
	"fmt"
	"net/http"
)

var (
	ErrUnauthorized          = fmt.Errorf("unauthorized")
	ErrUserAlreadyRegistered = fmt.Errorf("user already registered")
	ErrDataNotFound          = fmt.Errorf("data not found")
)

var mapErrToHTTPCode = map[error]int{
	ErrUnauthorized:          http.StatusUnauthorized,
	ErrUserAlreadyRegistered: http.StatusConflict,
	ErrDataNotFound:          http.StatusNotFound,
}

func ErrorToHTTPCode(err error) int {
	code, ok := mapErrToHTTPCode[err]
	if !ok {
		code = http.StatusInternalServerError
	}

	return code
}
