package twitch

import "fmt"

type ErrInvalidStatus struct {
	Status string
}

func (e ErrInvalidStatus) Error() string {
	return fmt.Sprintf("invalid HTTP status %s", e.Status)
}
