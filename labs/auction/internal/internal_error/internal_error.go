package internal_error

type InternalError struct {
	Message string
	Err     string
}

func (e *InternalError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "not_found",
	}
}

func NewInternalServerError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "internal_server",
	}
}
