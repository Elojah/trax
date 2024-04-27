package agg

import (
	"context"
	"strings"
	"time"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/pkg/cookie"
	"github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/elojah/trax/pkg/ulid"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/metadata"
)

var _ user.Agg = (*Agg)(nil)

type Agg struct {
	transaction.Transactioner

	user.Store

	user.StoreEntity

	user.StoreRole
	user.StorePermission
	user.StoreRoleUser

	Cookie cookie.Agg
}

func (a *Agg) Dial(ctx context.Context) error {
	return nil
}

func (a Agg) fetchClaimAuth(ctx context.Context, userID ulid.ID, audience string) (user.ClaimAuth, error) {
	var ca user.ClaimAuth

	// don't populate any claims for token other than access
	if audience != "access" {
		return ca, nil
	}

	if err := a.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		var err error

		ca, err = a.ListClaims(ctx, userID)
		if err != nil {
			return transaction.Rollback, err
		}

		return transaction.Commit, nil
	}); err != nil {
		return user.ClaimAuth{}, err
	}

	return ca, nil
}

func (a Agg) CreateJWT(ctx context.Context, u user.U, audience string, validity time.Duration) (string, error) {
	// #MARK:Create JWT
	id := ulid.NewID()

	// use cookie rotation encoding to generate a rotating secret for JWT
	secret, err := a.Cookie.Encode(ctx, "jwt_secret", string(id.Bytes()))
	if err != nil {
		return "", err
	}

	ca, err := a.fetchClaimAuth(ctx, u.ID, audience)
	if err != nil {
		return "", err
	}

	now := time.Now()
	claims := &user.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "trax",
			Subject:   u.ID.String(),
			Audience:  jwt.ClaimStrings{audience},
			ExpiresAt: jwt.NewNumericDate(now.Add(validity)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        secret,
		},
		UserID: u.ID,
		Auth:   ca,
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

func (a Agg) ReadJWT(ctx context.Context, token string) (user.Claims, error) {
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

func (a Agg) Auth(ctx context.Context, audience string) (user.Claims, error) {
	// read & parse token
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return user.Claims{}, errors.ErrMissingAuth{}
	}

	token := func() string {
		if len(md["token"]) == 0 {
			return ""
		}

		return strings.Trim(md["token"][0], " ")
	}()

	if token == "" {
		return user.Claims{}, errors.ErrMissingAuth{}
	}

	claims, err := a.ReadJWT(ctx, token)
	if err != nil {
		return user.Claims{}, err
	}

	if len(claims.Audience) == 0 || claims.Audience[0] != audience {
		return user.Claims{}, err
	}

	id, err := ulid.Parse(claims.Subject)
	if err != nil {
		return user.Claims{}, user.ErrInvalidClaims{Err: err}
	}

	claims.UserID = id

	return claims, nil
}
