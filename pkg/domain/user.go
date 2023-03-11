package domain

import (
	"context"

	"fsm/pkg/ent"
)

type UserRepository interface {
	Create(ctx context.Context, user ent.User) error
	GetByEmail(ctx context.Context, email string) (*ent.User, error)
	GetByID(ctx context.Context, uid string) (*ent.User, error)
	UpdatePassword(ctx context.Context, user ent.User) error
}
