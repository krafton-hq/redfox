package errors

import "fmt"

type Error interface {
	error
	Message() string
	Origin() error
}

type redFoxErrorFundamental struct {
	message string
	origin  error
}

func (e *redFoxErrorFundamental) Error() string {
	if e.origin == nil {
		return e.message
	} else {
		return e.message + ", error: " + e.origin.Error()
	}
}

func (e *redFoxErrorFundamental) Message() string {
	return e.message
}

func (e *redFoxErrorFundamental) Origin() error {
	return e.origin
}

func NewError(message string) Error {
	return newError(nil, message)
}

func NewErrorf(format string, args ...any) Error {
	return newError(nil, fmt.Sprintf(format, args...))
}

func WrapError(origin error, message string) Error {
	return newError(origin, message)
}

func WrapErrorf(origin error, format string, args ...any) Error {
	return newError(origin, fmt.Sprintf(format, args...))
}

func newError(origin error, message string) Error {
	return &redFoxErrorFundamental{
		message: message,
		origin:  origin,
	}
}
