package migrate

import "context"

type Agg interface {
	Up(context.Context, string) error
	Down(context.Context, string) error
}
