package google

import "fmt"

type ErrInvalidClaim struct {
	Claim string
}

func (e ErrInvalidClaim) Error() string {
	return fmt.Sprintf("invalid claim %s", e.Claim)
}
