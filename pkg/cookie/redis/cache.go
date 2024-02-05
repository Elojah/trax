package redis

import (
	"github.com/elojah/trax/pkg/cookie"
	"github.com/elojah/trax/pkg/redis"
)

var _ cookie.CacheKeys = (*Cache)(nil)

type Cache struct {
	redis.Service
}
