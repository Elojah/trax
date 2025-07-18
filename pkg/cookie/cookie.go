package cookie

import "context"

const (
	keyLength = 64
	NKeys     = 5
)

// Service

type Service interface {
	CacheKeys

	Encode(context.Context, string, string) (string, error)
	Decode(context.Context, string, string) (string, error)
}
