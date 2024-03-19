package errors

import "fmt"

type ErrUnknown struct{}

func (err ErrUnknown) Error() string {
	return fmt.Sprintf("unknown error, investigation required")
}

type ErrNotImplemented struct {
	Version string
}

func (err ErrNotImplemented) Error() string {
	return fmt.Sprintf("not implemented yet, planned for %s", err.Version)
}

type ErrMissingImplementation struct {
	Interface string
}

func (err ErrMissingImplementation) Error() string {
	return fmt.Sprintf("interface %s must be set", err.Interface)
}
