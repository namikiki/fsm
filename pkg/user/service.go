package user

import (
	"context"
	"errors"
	"fsm/pkg/domain"
	"fsm/pkg/ent"
	"fsm/pkg/jwt"
	"fsm/pkg/salt"
	"github.com/google/uuid"
	"log"
)

type Service struct {
	jwt  jwt.Service
	salt salt.Service
	user domain.UserRepository
}

func NewService(jwt jwt.Service, salt salt.Service, user domain.UserRepository) domain.UserService {
	return &Service{
		jwt:  jwt,
		salt: salt,
		user: user,
	}
}

func (s *Service) Login(ctx context.Context, email, reqPassword string) (*ent.User, error) {
	user, err := s.user.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	comp := s.ComparePassword(reqPassword, user.Salt, user.PassWord)
	if !comp {
		return nil, errors.New("password error")
	}

	//token, err := s.jwt.Gen(ctx, user.ID)
	//if err != nil {
	//	return nil, "", err
	//}

	return user, nil
}

func (s *Service) ComparePassword(reqPassword string, salt []byte, hashedPassword string) bool {
	hashed := s.salt.Hashed([]byte(reqPassword), salt)
	return hashed == hashedPassword
}

func (s *Service) Register(ctx context.Context, email, password, username string) (*ent.User, error) {
	if _, err := s.user.GetByEmail(ctx, email); err == nil {
		return nil, errors.New("用户已注册")
	}

	saltStr := s.salt.RandBytesSlice(len(password))
	hashed := s.salt.Hashed([]byte(password), saltStr)
	uid := uuid.New().String()

	u := ent.User{
		ID:              uid,
		Email:           email,
		PassWord:        hashed,
		Salt:            saltStr,
		UserName:        username,
		BucketName:      uid,
		CurrentStoreCap: 0,
		MaxStoreCap:     1000000,
	}

	if err := s.user.Store(ctx, u); err != nil {
		log.Printf("注册失败 %v", err)
		return nil, err
	}

	//token, err := s.jwt.Gen(ctx, u.ID)
	//if err != nil {
	//	return nil, errors.New("生成token失败")
	//}
	return &u, nil
}
