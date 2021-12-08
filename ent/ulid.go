package ent

import (
	"context"
	"fmt"
	"golang-clean-architecture-ent-gqlgen/ent/schema/ulid"
	"golang-clean-architecture-ent-gqlgen/pkg/const/globalid"
)

var globalIDs = globalid.New()

// IDToType maps an ulid.ID to the underlying table.
func IDToType(_ context.Context, id ulid.ID) (string, error) {
	if len(id) < 3 {
		return "", fmt.Errorf("IDToType: id too short")
	}
	prefix := id[:3]
	t, err := globalIDs.FindTableByID(string(prefix))
	if err != nil {
		return "", fmt.Errorf("IDToType: could not map prefix '%s' to a type", prefix)
	}
	return t, nil
}
