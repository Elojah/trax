package user

import (
	"github.com/elojah/trax/pkg/ulid"
	"github.com/golang-jwt/jwt/v4"
)

type Requirement struct {
	GroupID  ulid.ID
	Resource Resource
	Command  Command
}

func NewRequirements(groupIDs []ulid.ID, resource Resource, command Command) []Requirement {
	requirements := make([]Requirement, 0, len(groupIDs))
	for _, id := range groupIDs {
		requirements = append(requirements, Requirement{
			GroupID:  id,
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

func (c ClaimGroup) Require(rs ...Requirement) error {
	for _, r := range rs {
		resources, ok := c.Resources[r.Resource.String()]
		if !ok {
			return ErrUnauthorizedResource{Resource: r.Resource.String()}
		}

		if _, ok := resources.Commands[r.Command.String()]; !ok {
			return ErrUnauthorizedCommand{Resource: r.Resource.String(), Command: r.Command.String()}
		}
	}

	return nil
}

func (c Claims) GroupIDs(rs ...Requirement) []ulid.ID {
	result := make([]ulid.ID, 0, len(c.Auth.Groups))
	for id, group := range c.Auth.Groups {
		groupID, err := ulid.Parse(id)
		if err != nil {
			continue
		}

		if err := group.Require(rs...); err != nil {
			// ignore error but don't add groupID to result
			continue
		}

		result = append(result, groupID)
	}

	return result
}

func (c Claims) RoleIDs() []ulid.ID {
	result := make([]ulid.ID, 0, len(c.Auth.Groups))
	for _, group := range c.Auth.Groups {
		for roleID := range group.Roles {
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
		groups, ok := c.Auth.Groups[r.GroupID.String()]
		if !ok {
			return ErrUnknownGroup{GroupID: r.GroupID.String()}
		}

		resources, ok := groups.Resources[r.Resource.String()]
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

	for _, g := range c.Auth.Groups {
		for id := range g.Roles {
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
