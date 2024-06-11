package repositories

import (
	"context"
	"fsm/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(ctx context.Context, user models.User) error {
	u.db.Create(&user)
	return nil
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	u.db.Where("email = ?", email).Find(&user)
	return &user, nil
}

func (u *UserRepository) GetByID(ctx context.Context, uid string) (*models.User, error) {
	var user models.User
	u.db.Where("id = ?", uid).Find(&user)
	return &user, nil
}

func (u *UserRepository) UpdatePassword(ctx context.Context, user models.User) error {
	u.db.Save(&user)
	return nil
}
