package services

import (
	"context"
	"errors"
	"fsm/models"
	"fsm/pkg/repositories"
	"fsm/pkg/utils"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"time"
)

const (
	LoginAttemptLimit  = 5
	LoginAttemptWindow = 15 * time.Minute
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

// UserService 用户服务
type UserService struct {
	userRepo *repositories.UserRepository
	jwt      *JWTService
	redis    *redis.Client
}

// NewUserService 创建一个新的 UserService 实例
func NewUserService(jwt *JWTService, redis *redis.Client, userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		jwt:      jwt,
		userRepo: userRepo,
		redis:    redis,
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
	res := s.redis.Get(ctx, uls.Email)
	if res.Err() != nil {
		if errors.Is(res.Err(), redis.Nil) {
			return "", errors.New("key does not exist")
		}
		return "", res.Err()
	}
	if intVal, _ := res.Int(); intVal >= 5 {
		return "", errors.New("每日登录重试次数超过")
	}
	user, err := s.userRepo.GetByEmail(ctx, uls.Email)

	//校验密码和账号状态
	if err != nil {
		return "", errors.New("用户或者邮箱错误")
	}
	if !user.CheckPassword(uls.PassWord) {
		return "", errors.New("密码错误")
	}
	if user.Status == "baned" {
		return "", errors.New("账户被封禁")
	}

	////todo 存储
	//clientID := md5.Sum([]byte(uls.MAC))

	//生成并返回JWT
	token, err := s.jwt.Gen(ctx, user.ID)
	if err != nil {
		//todo log
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

// GetUser 根据用户ID获取用户信息
func (s *UserService) GetUser(ctx context.Context, userID string) (*models.User, error) {
	return s.userRepo.GetByID(ctx, userID)
}

// CheckLoginAttempts 检查用户登录尝试次数
func (s *UserService) CheckLoginAttempts(ctx context.Context, username string) (bool, error) {
	key := "login_attempts:" + username
	attempts, err := s.redis.Get(ctx, key).Int()

	if err != nil && !errors.Is(err, redis.Nil) {
		return false, err
	}

	if attempts >= LoginAttemptLimit {
		return false, nil
	}

	return true, nil
}

// IncrementLoginAttempts 增加用户尝试次数
func (s *UserService) IncrementLoginAttempts(ctx context.Context, username string) error {
	key := "login_attempts:" + username
	_, err := s.redis.Incr(ctx, key).Result()
	if err != nil {
		return err
	}
	s.redis.Expire(ctx, key, LoginAttemptWindow)
	return nil
}

// ResetLoginAttempts  重置用户登录尝试次数
func (s *UserService) ResetLoginAttempts(ctx context.Context, username string) error {
	key := "login_attempts:" + username
	return s.redis.Del(ctx, key).Err()
}
