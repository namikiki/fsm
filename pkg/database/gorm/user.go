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

func (u *UserRepository) Create(ctx context.Context, user ent.User) error {
	u.Conn.Create(&user)
	return nil
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (ent.User, error) {
	var user ent.User
	u.Conn.Where("email = ?", email).Find(&user)
	return user, nil
}

func (u *UserRepository) GetByID(ctx context.Context, uid string) (ent.User, error) {
	var user ent.User
	u.Conn.Where("id = ?", uid).Find(&user)
	return user, nil
}
