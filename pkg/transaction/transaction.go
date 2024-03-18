package transaction

import "context"

type Key struct{}

type AccessMode int

const (
	Read AccessMode = iota
	Write
)

type Operation int

const (
	Rollback Operation = iota
	Commit
)

type Transactioner interface {
	Tx(context.Context, AccessMode, func(context.Context) (Operation, error)) error
}
