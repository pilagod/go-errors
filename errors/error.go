package errors

import (
	"fmt"
)

// Error an error with code, message, and other configurable data
type Error struct {
	code           string
	data           interface{}
	message        string
	httpStatusCode *int
}

// Error returns error message in format "[code] message"
func (e *Error) Error() string {
	return fmt.Sprintf("[%s] %s", e.code, e.message)
}

// Code returns code on error
func (e *Error) Code() string {
	return e.code
}

// Data returns data on error
func (e *Error) Data() interface{} {
	return e.data
}

// Message returns message on error
func (e *Error) Message() string {
	return e.message
}

// HTTPStatusCode returns http status code on error (default: 500)
func (e *Error) HTTPStatusCode() int {
	if e.httpStatusCode == nil {
		return 500
	}
	return *e.httpStatusCode
}
