package service

import (
	"context"

	"github.com/elojah/trax/pkg/cookie"
	"github.com/gorilla/securecookie"
)

const maxCookieLength = 131072

type S struct {
	cookie.CacheKeys
}

// Initial keys setup

func (s *S) Setup(ctx context.Context, n int) error {
	keys := make([]cookie.Keys, 0, n)

	if err := s.DeleteKeys(ctx, cookie.FilterKeys{
		All: true,
	}); err != nil {
		return err
	}

	for i := 0; i < n; i++ {
		keys = append(keys, cookie.NewKeys())
	}

	if err := s.CreateKeys(ctx, keys...); err != nil {
		return err
	}

	return nil
}

// Encoding/Decoding helpers

func (s S) Encode(ctx context.Context, key string, value string) (string, error) {
	ks, err := s.ReadKeys(ctx, cookie.FilterKeys{All: true})
	if err != nil {
		return "", err
	}

	if len(ks) == 0 {
		return "", cookie.ErrMissingSecureKeys{}
	}

	ck, err := securecookie.New(
		ks[len(ks)-1].Hash,
		ks[len(ks)-1].Block,
	).MaxLength(maxCookieLength).Encode(key, value)
	if err != nil {
		return "", err
	}

	return ck, nil
}

func (s S) Decode(ctx context.Context, key string, value string) (string, error) {
	keys, err := s.ReadKeys(ctx, cookie.FilterKeys{All: true})
	if err != nil {
		return "", err
	}

	scs := func() []securecookie.Codec {
		result := make([]securecookie.Codec, 0, len(keys))
		for _, k := range keys {
			result = append(result, securecookie.New(k.Hash, k.Block).MaxLength(maxCookieLength))
		}

		return result
	}()

	var result string

	if err := securecookie.DecodeMulti(key, value, &s, scs...); err != nil {
		return "", err
	}

	return result, nil
}
