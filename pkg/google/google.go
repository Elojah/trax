package google

import (
	"context"
	"time"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/pkg/ulid"
	"golang.org/x/oauth2"
)

type Claims map[string]any

type App interface {
	Signin(context.Context, string) (string, Claims, error)
	OAuth() *oauth2.Config
}

func (c Claims) GetString(key string) (string, error) {
	// #Retrieve Google ID
	val, ok := c[key]
	if !ok {
		return "", ErrInvalidClaim{Claim: "claim not found"}
	}

	result, ok := val.(string)
	if !ok {
		return "", ErrInvalidClaim{Claim: "invalid claim format"}
	}

	return result, nil
}

func (c Claims) CreateUser(ctx context.Context) (user.U, user.Profile, error) {
	gid, err := c.GetString("sub")
	if err != nil {
		return user.U{}, user.Profile{}, err
	}

	email, err := c.GetString("email")
	if err != nil {
		return user.U{}, user.Profile{}, err
	}

	lastname, err := c.GetString("family_name")
	if err != nil {
		return user.U{}, user.Profile{}, err
	}

	firstname, err := c.GetString("given_name")
	if err != nil {
		return user.U{}, user.Profile{}, err
	}

	avatarURL, err := c.GetString("picture")
	if err != nil {
		return user.U{}, user.Profile{}, err
	}

	// #Create user
	now := time.Now().Unix()
	// generate a random password for google sign
	hash, salt := user.Encrypt(ulid.NewID().String())

	u := user.U{
		ID:           ulid.NewID(),
		Email:        email,
		PasswordHash: hash,
		PasswordSalt: salt,
		GoogleID:     gid,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	profile := user.Profile{
		UserID:    u.ID,
		FirstName: firstname,
		LastName:  lastname,
		AvatarURL: avatarURL,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return u, profile, nil
}
