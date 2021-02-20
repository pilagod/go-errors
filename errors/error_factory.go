package errors

import "github.com/pkg/errors"

// ErrorFactory factory for error
type ErrorFactory func(message string, options ...Option) error

// Create returns an error factory which creates error with code and options
func Create(code string, defaultOptions ...Option) ErrorFactory {
	return func(message string, overriddenOptions ...Option) error {
		err := &Error{
			code:    code,
			message: message,
		}
		for _, option := range append(defaultOptions, overriddenOptions...) {
			option(err)
		}
		return errors.WithStack(err)
	}
}

// Option for error
type Option func(e *Error)

// Data sets data on error
func Data(data interface{}) Option {
	return func(e *Error) {
		e.data = data
	}
}

// HTTPStatusCode sets http status code on error
func HTTPStatusCode(statusCode int) Option {
	return func(e *Error) {
		e.httpStatusCode = &statusCode
	}
}
