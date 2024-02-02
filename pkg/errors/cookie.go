package errors

import "fmt"

type ErrMissingSecureKeys struct{}

func (err ErrMissingSecureKeys) Error() string {
	return fmt.Sprintf("cookie keys not found")
}
