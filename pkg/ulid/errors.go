package ulid

import "fmt"

type ErrInvalidULID struct {
	Value any
}

func (e ErrInvalidULID) Error() string {
	return fmt.Sprintf("invalid ulid %v", e.Value)
}
