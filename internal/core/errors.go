package core

type InvariantViolationError struct {
	Message string
}

func (e *InvariantViolationError) Error() string {
	return e.Message
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

type PermissionDeniedError struct {
	Message string
}

func (e *PermissionDeniedError) Error() string {
	return e.Message
}
