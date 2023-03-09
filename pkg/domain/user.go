package domain

import (
	"context"

	"fsm/pkg/ent"
)

type UserService interface {
	Login(ctx context.Context, email, reqPassword string) (*ent.User, error)
	ComparePassword(reqPassword string, salt []byte, hashedPassword string) bool
	Register(ctx context.Context, email, password, username string) (*ent.User, error)
}

type UserRepository interface {
	Store(ctx context.Context, u ent.User) error
	GetByEmail(ctx context.Context, email string) (*ent.User, error)
	GetByID(ctx context.Context, uid string) (*ent.User, error)
}
