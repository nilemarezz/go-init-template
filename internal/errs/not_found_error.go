package errs

// NotFoundError represents an error when a resource is not found.
type NotFoundError struct {
	Resource string
}

// NewNotFoundError creates a new NotFoundError.
func NewNotFoundError(resource string) error {
	return &NotFoundError{Resource: resource}
}

// Error returns the error message for NotFoundError.
func (e NotFoundError) Error() string {
	return e.Resource + " not found"
}
