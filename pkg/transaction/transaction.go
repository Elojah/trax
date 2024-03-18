package transaction

import "context"

type Key struct{}

type AccessMode int

const (
	NilAccessMode AccessMode = iota
	Read
	Write
)

type Operation int

const (
	NilOperation Operation = iota
	Rollback
	Commit
)

type Transactioner interface {
	Tx(context.Context, AccessMode, func(context.Context) (Operation, error)) error
}
