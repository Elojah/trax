package migrate

import "context"

type App interface {
	Up(context.Context, string) error
	Down(context.Context, string) error
}
