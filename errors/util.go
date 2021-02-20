package errors

import (
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// Cause returns error under stacktracer
func Cause(err error) error {
	return errors.Cause(err)
}

// Is checks err is same type as errRef
func Is(err error, errRef interface{}) bool {
	e, ok := Cause(err).(*Error)
	if !ok {
		return false
	}
	var code string
	switch t := errRef.(type) {
	case string:
		code = t
	case ErrorFactory:
		code = Cause(t("")).(*Error).Code()
	}
	return e.code == code
}

// StackTrace returns error stacktrace
func StackTrace(err error) errors.StackTrace {
	e, ok := err.(stackTracer)
	if !ok {
		return nil
	}
	return e.StackTrace()[1:]
}
