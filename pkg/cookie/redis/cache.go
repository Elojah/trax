package redis

import (
	"github.com/elojah/redis"
	"github.com/elojah/trax/pkg/cookie"
)

var _ cookie.CacheKeys = (*Cache)(nil)

type Cache struct {
	redis.Service
}
