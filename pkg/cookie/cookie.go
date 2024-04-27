package cookie

import "context"

const (
	keyLength = 64
	NKeys     = 5
)

// Agg

type Agg interface {
	CacheKeys

	Encode(context.Context, string, string) (string, error)
	Decode(context.Context, string, string) (string, error)
}
