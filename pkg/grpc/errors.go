package grpc

// ErrPermissionForbidden is raised when a user tries to access a forbidden resource.
type ErrPermissionForbidden struct{}

// Error implementation for ErrPermissionForbidden.
func (e ErrPermissionForbidden) Error() string {
	return "forbidden access"
}

// ErrUninitialized is raised when grpc service serve after wrong/no dial.
type ErrUninitialized struct{}

// Error implementation for ErrUninitialized.
func (err ErrUninitialized) Error() string {
	return "grpc service is not initialized"
}
