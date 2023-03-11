package user

import (
	"context"
	"errors"
	"log"

	"fsm/api/req"
	"fsm/pkg/domain"
	"fsm/pkg/ent"
	"fsm/pkg/jwt"
	"fsm/pkg/salt"

	"github.com/google/uuid"
)

type Service struct {
	jwt  jwt.Service
	salt salt.Service
	user domain.UserRepository
}

func NewService(jwt jwt.Service, salt salt.Service, user domain.UserRepository) *Service {
	return &Service{
		jwt:  jwt,
		salt: salt,
		user: user,
	}
}

func (s *Service) Login(ctx context.Context, email, password string) (string, string, error) {
	user, err := s.user.GetByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	comp := s.ComparePassword(password, user.Salt, user.PassWord)
	if !comp {
		return "", "", errors.New("password error")
	}

	token, err := s.jwt.Gen(ctx, user.ID)
	if err != nil {
		return "", "", err
	}

	return user.ID, token, nil
}

func (s *Service) UpdatePassword(ctx context.Context, up req.UpdatePassword) error {
	user, err := s.user.GetByEmail(ctx, up.Email)
	if err != nil {
		return err
	}

	if comp := s.ComparePassword(up.OldPassword, user.Salt, user.PassWord); !comp {
		return errors.New("password error")
	}

	user.Salt = s.salt.RandBytesSlice(len(user.PassWord))
	user.PassWord = s.salt.Hashed([]byte(user.PassWord), user.Salt)

	return s.user.UpdatePassword(ctx, *user)

}

func (s *Service) ComparePassword(reqPassword string, salt []byte, hashedPassword string) bool {
	hashed := s.salt.Hashed([]byte(reqPassword), salt)
	return hashed == hashedPassword
}

func (s *Service) Register(ctx context.Context, user req.UserRegister) (*ent.User, error) {
	if u, _ := s.user.GetByEmail(ctx, user.Email); u.ID != "" {
		return nil, errors.New("用户已注册")
	}

	saltStr := s.salt.RandBytesSlice(len(user.PassWord))
	hashed := s.salt.Hashed([]byte(user.PassWord), saltStr)
	uid := uuid.New().String()

	u := ent.User{
		ID:              uid,
		Email:           user.Email,
		PassWord:        hashed,
		Salt:            saltStr,
		UserName:        user.UserName,
		BucketName:      uid,
		CurrentStoreCap: 0,
		MaxStoreCap:     1000000,
	}

	if err := s.user.Create(ctx, u); err != nil {
		log.Printf("注册失败 %v", err)
		return nil, err
	}
	return &u, nil
}
