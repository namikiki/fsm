package ent

import (
	"context"

	"fsm/pkg/domain"
	"fsm/pkg/ent"
	"fsm/pkg/ent/user"
)

type mysqlUserRepository struct {
	Conn *ent.Client
}

func NewMysqlUserRepository(Conn *ent.Client) domain.UserRepository {
	return &mysqlUserRepository{Conn: Conn}
}

func (s *mysqlUserRepository) Store(ctx context.Context, u ent.User) error {
	_, err := s.Conn.User.Create().
		SetID(u.ID).
		SetUserName(u.UserName).
		SetPassWord(u.PassWord).
		SetEmail(u.Email).
		SetSalt(u.Salt).
		SetBucketName(u.BucketName).
		SetMaxStoreCap(u.MaxStoreCap).
		SetCurrentStoreCap(u.CurrentStoreCap).
		Save(ctx)
	return err
}

func (s *mysqlUserRepository) GetByEmail(ctx context.Context, email string) (*ent.User, error) {
	return s.Conn.User.Query().
		Where(user.Email(email)).Only(ctx)
}

func (s *mysqlUserRepository) GetByID(ctx context.Context, uid string) (*ent.User, error) {
	return s.Conn.User.Query().
		Where(user.ID(uid)).Only(ctx)
}

//func (s *sqliteUserRepository) Fetch(ctx context.Context) ([]*ent.User, error) {
//	return s.Conn.User.Query().All(ctx)
//}
//
//func (s *sqliteUserRepository) GetByID(ctx context.Context, id string) (*ent.User, error) {
//	return s.Conn.User.Query().Where(user.ID(id)).Only(ctx)
//}
