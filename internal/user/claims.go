package user

import (
	"github.com/elojah/trax/pkg/ulid"
	"github.com/golang-jwt/jwt/v4"
)

type Requirement struct {
	EntityID ulid.ID
	Resource Resource
	Command  Command
}

func NewRequirements(entityIDs []ulid.ID, resource Resource, command Command) []Requirement {
	requirements := make([]Requirement, 0, len(entityIDs))
	for _, id := range entityIDs {
		requirements = append(requirements, Requirement{
			EntityID: id,
			Resource: resource,
			Command:  command,
		})
	}

	return requirements
}

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

func (c Claims) RoleIDs() []ulid.ID {
	result := make([]ulid.ID, 0, len(c.Auth.Entities))
	for _, entity := range c.Auth.Entities {
		for roleID := range entity.Roles {
			id, err := ulid.Parse(roleID)
			if err != nil {
				continue
			}
			result = append(result, id)
		}
	}

	return result
}

func (c Claims) Require(rs ...Requirement) error {
	for _, r := range rs {
		entities, ok := c.Auth.Entities[r.EntityID.String()]
		if !ok {
			return ErrUnknownEntity{EntityID: r.EntityID.String()}
		}

		resources, ok := entities.Resources[r.Resource.String()]
		if !ok {
			return ErrUnauthorizedResource{Resource: r.Resource.String()}
		}

		if _, ok := resources.Commands[r.Command.String()]; !ok {
			return ErrUnauthorizedCommand{Resource: r.Resource.String(), Command: r.Command.String()}
		}
	}

	return nil
}

func (c Claims) RequireRoles(ids ...ulid.ID) error {
	mapIDs := make(map[string]struct{}, len(ids))
	for _, id := range ids {
		mapIDs[id.String()] = struct{}{}
	}

	for _, e := range c.Auth.Entities {
		for id := range e.Roles {
			if _, ok := mapIDs[id]; ok {
				delete(mapIDs, id)
				if len(mapIDs) == 0 {
					return nil
				}
			}
		}
	}

	missingRoleIDs := make([]string, 0, len(mapIDs))
	for id := range mapIDs {
		missingRoleIDs = append(missingRoleIDs, id)
	}

	return ErrUnauthorizedRole{Roles: missingRoleIDs}
}
