package errors

type NotFoundError struct {
	err Error
}

func (e *NotFoundError) Error() string {
	return e.err.Error()
}

func (e *NotFoundError) Message() string {
	return e.err.Message()
}

func (e *NotFoundError) Origin() error {
	return e.err.Origin()
}

const notFoundMessage = "Resource Not Found Id: %s"

func NewNotFound(id string) Error {
	return WrapNotFound(nil, id)
}

func WrapNotFound(origin error, id string) Error {
	return &NotFoundError{err: WrapErrorf(origin, notFoundMessage, id)}
}

type InvalidArgumentsError struct {
	err Error
}

func (e *InvalidArgumentsError) Error() string {
	return e.err.Error()
}

func (e *InvalidArgumentsError) Message() string {
	return e.err.Message()
}

func (e *InvalidArgumentsError) Origin() error {
	return e.err.Origin()
}

const invalidArguments = "Invalid Argument Rule: %s"

func NewInvalidArguments(rule string) Error {
	return WrapInvalidArguments(nil, rule)
}

func WrapInvalidArguments(origin error, id string) Error {
	return &InvalidArgumentsError{err: WrapErrorf(origin, invalidArguments, id)}
}

type InternalError struct {
	err Error
}

func (e *InternalError) Error() string {
	return e.err.Error()
}

func (e *InternalError) Message() string {
	return e.err.Message()
}

func (e *InternalError) Origin() error {
	return e.err.Origin()
}

const internalError = "Unexpected Internal Error Occurred: %s"

func NewInternalError(detail string) Error {
	return WrapInternalError(nil, detail)
}

func WrapInternalError(origin error, detail string) Error {
	return &InternalError{err: WrapErrorf(origin, internalError, detail)}
}
