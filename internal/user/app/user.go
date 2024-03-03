package app

import (
	"context"
	"strings"
	"time"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/pkg/cookie"
	"github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/ulid"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/metadata"
)

var _ user.App = (*App)(nil)

type App struct {
	user.Store
	user.StoreProfile

	Cookie cookie.App
}

func (a *App) Dial(ctx context.Context) error {
	return nil
}

func (a App) CreateJWT(ctx context.Context, u user.U, audience string, validity time.Duration) (string, error) {
	// #Create JWT
	id := ulid.NewID()

	// use cookie rotation encoding to generate a rotating secret for JWT
	secret, err := a.Cookie.Encode(ctx, "jwt_secret", string(id.Bytes()))
	if err != nil {
		return "", err
	}

	now := time.Now()
	claims := &user.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "spc",
			Subject:   u.ID.String(),
			Audience:  jwt.ClaimStrings{audience},
			ExpiresAt: jwt.NewNumericDate(now.Add(validity)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        secret,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(id.Bytes())
	if err != nil {
		return "", err
	}

	ck, err := a.Cookie.Encode(ctx, "token", ss)
	if err != nil {
		return "", err
	}

	return ck, nil
}

func (a App) ReadJWT(ctx context.Context, token string) (user.Claims, error) {
	ck, err := a.Cookie.Decode(ctx, "token", token)
	if err != nil {
		return user.Claims{}, err
	}

	t, err := jwt.ParseWithClaims(ck, &user.Claims{}, func(t *jwt.Token) (any, error) {
		claims, ok := t.Claims.(*user.Claims)
		if !ok {
			return user.Claims{}, user.ErrInvalidClaims{}
		}

		secret, err := a.Cookie.Decode(ctx, "jwt_secret", claims.RegisteredClaims.ID)
		if err != nil {
			return "", err
		}

		return []byte(secret), nil
	})
	if err != nil {
		return user.Claims{}, err
	}

	// check token validity
	claims, ok := t.Claims.(*user.Claims)
	if !ok {
		return user.Claims{}, user.ErrInvalidClaims{}
	}

	if err := claims.Valid(); err != nil {
		return user.Claims{}, user.ErrInvalidClaims{Err: err}
	}

	return *claims, nil
}

func (a App) Auth(ctx context.Context, audience string) (user.U, error) {
	// read & parse token
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return user.U{}, errors.ErrMissingAuth{}
	}

	token := func() string {
		if len(md["token"]) == 0 {
			return ""
		}

		return strings.Trim(md["token"][0], " ")
	}()

	if token == "" {
		return user.U{}, errors.ErrMissingAuth{}
	}

	claims, err := a.ReadJWT(ctx, token)
	if err != nil {
		return user.U{}, err
	}

	if len(claims.Audience) == 0 || claims.Audience[0] != audience {
		return user.U{}, err
	}

	id, err := ulid.Parse(claims.Subject)
	if err != nil {
		return user.U{}, user.ErrInvalidClaims{Err: err}
	}

	return user.U{
		ID: id,
	}, nil
}
