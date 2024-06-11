package services

import (
	"context"
	"errors"
	"fsm/models"
	"fsm/pkg/repositories"
	"fsm/pkg/utils"
	"github.com/google/uuid"
)

// UserLoginService  用户登录服务
type UserLoginService struct {
	Email    string `json:"email" validate:"required"`
	PassWord string `json:"password" validate:"required"`
	MAC      string `json:"mac" validate:"required"`
}

// UserRegisterService 用户注册服务
type UserRegisterService struct {
	Email    string `json:"email,omitempty" validate:"required,min=10,email"`
	PassWord string `json:"password,omitempty" validate:"required,min=10"`
	UserName string `json:"username,omitempty" validate:"required,min=5"`
}

// UpdatePasswordService 密码更新服务
type UpdatePasswordService struct {
	Email       string `json:"email" validate:"required"`
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type UserService struct {
	userRepo *repositories.UserRepository
	jwt      *JWTService
}

func NewUserService(jwt *JWTService, userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		jwt:      jwt,
		userRepo: userRepo,
	}
}

// Register  新用户注册
func (s *UserService) Register(ctx context.Context, urs UserRegisterService) (*models.User, error) {
	if u, err := s.userRepo.GetByEmail(ctx, urs.Email); u.ID != "" || err != nil {
		return nil, errors.Join(err, errors.New("用户已注册"))
	}

	uid := uuid.New().String()
	salt := utils.RandBytesSlice(len(urs.PassWord))
	passwordHash := utils.PasswordHash(urs.PassWord, salt)

	//todo new User
	u := models.User{
		ID:              uid,
		UserName:        urs.UserName,
		Email:           urs.Email,
		PassWord:        passwordHash,
		Salt:            salt,
		BucketName:      uid,
		CurrentStoreCap: 0,
		MaxStoreCap:     1000000,
	}

	return &u, s.userRepo.Create(ctx, &u)
}

// Login 用户密码登录
func (s *UserService) Login(ctx context.Context, uls UserLoginService) (string, error) {
	user, err := s.userRepo.GetByEmail(ctx, uls.Email)
	if err != nil {
		return "", err
	}

	if !user.CheckPassword(uls.PassWord) {
		return "", errors.New("密码错误")
	}

	////todo 存储
	//clientID := md5.Sum([]byte(uls.MAC))

	token, err := s.jwt.Gen(ctx, user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// UpdatePassword 用户更新密码
func (s *UserService) UpdatePassword(ctx context.Context, ups UpdatePasswordService) error {
	user, err := s.userRepo.GetByEmail(ctx, ups.Email)
	if err != nil {
		return err
	}

	if !user.CheckPassword(ups.OldPassword) {
		return errors.New("密码错误")
	}

	user.PassWord = utils.PasswordHash(ups.NewPassword, user.Salt)

	return s.userRepo.UpdatePassword(ctx, *user)
}

func (s *UserService) GetUser(ctx context.Context, userID string) (*models.User, error) {
	return s.userRepo.GetByID(ctx, userID)
}
