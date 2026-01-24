package repository

import (
	"context"
	"nguyendangviet12022004/url-short/user/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{Db: db}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	return gorm.G[model.User](ur.Db).Create(ctx, user)
}

func (ur *userRepository) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := gorm.G[model.User](ur.Db).Where("email = ?", email).First(ctx)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
