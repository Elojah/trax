package user

import (
	"github.com/elojah/trax/pkg/ulid"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims

	UserID ulid.ID
	Auth   ClaimAuth
}

func (c Claims) EntityIDs() []ulid.ID {
	result := make([]ulid.ID, 0, len(c.Auth.Entities))
	for entityKey := range c.Auth.Entities {
		entityID, err := ulid.Parse(entityKey)
		if err != nil {
			continue
		}

		result = append(result, entityID)
	}

	return result
}

func (c Claims) Require(entityID ulid.ID, resource Resource, command Command) error {
	entities, ok := c.Auth.Entities[entityID.String()]
	if !ok {
		return ErrUnknownEntity{EntityID: entityID.String()}
	}

	commands, ok := entities.Commands[resource.String()]
	if !ok {
		return ErrUnauthorizedResource{Resource: resource.String()}
	}

	for _, c := range commands.Commands {
		if command == c {
			return nil
		}
	}

	return ErrUnauthorizedCommand{Resource: resource.String(), Command: command.String()}
}
