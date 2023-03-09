package gorm

import (
	"context"

	"fsm/pkg/domain"
	"fsm/pkg/ent"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) domain.UserRepository {
	return &UserRepository{Conn: conn}
}

func (u *UserRepository) Store(ctx context.Context, user ent.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (*ent.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) GetByID(ctx context.Context, uid string) (*ent.User, error) {
	//TODO implement me
	panic("implement me")
}
