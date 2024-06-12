package repositories

import (
	"context"
	"fsm/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// UserRepository
// todo 缓存
type UserRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

// NewUserRepository 使用提供的 gorm.DB 和 redis.Client 实例创建一个新的 UserRepository 实例。
func NewUserRepository(db *gorm.DB, redis *redis.Client) *UserRepository {
	return &UserRepository{db: db, redis: redis}
}

// Create 在数据库中创建一个新的用户记录。
func (u *UserRepository) Create(ctx context.Context, user *models.User) error {
	return u.db.Create(user).Error
}

// GetByEmail 根据邮箱检索用户。
func (u *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	u.db.Where("email = ?", email).Find(&user)
	return &user, nil
}

// GetByID 根据用户ID检索用户。
func (u *UserRepository) GetByID(ctx context.Context, uid string) (*models.User, error) {
	var user models.User
	u.db.Where("id = ?", uid).Find(&user)
	return &user, nil
}

// UpdatePassword 更新用户密码。
func (u *UserRepository) UpdatePassword(ctx context.Context, user models.User) error {
	return u.db.Save(&user).Error
}
