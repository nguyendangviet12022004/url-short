package service

import (
	"context"
	"nguyendangviet12022004/url-short/user/internal/dto"
	"nguyendangviet12022004/url-short/user/internal/exception"
	"nguyendangviet12022004/url-short/user/internal/model"
	"nguyendangviet12022004/url-short/user/internal/repository"
	"nguyendangviet12022004/url-short/user/pkg/jwt"
	"nguyendangviet12022004/url-short/user/pkg/logging"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type UserService interface {
	RegisterUser(ctx context.Context, req *dto.RegisterRequest) error
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (us *userService) RegisterUser(ctx context.Context, req *dto.RegisterRequest) error {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user := &model.User{
		Email:    req.Email,
		Password: string(hashPassword),
	}

	existUser, _ := us.userRepo.FindUserByEmail(ctx, req.Email)

	if existUser != nil {
		logging.GetLogger().Error(user.Email)
		return exception.ErrUserExists
	}

	return us.userRepo.CreateUser(ctx, user)
}

func (us *userService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := us.userRepo.FindUserByEmail(ctx, req.Email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, exception.ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, exception.ErrBadCredential
	}

	accessToken, err := jwt.GenerateToken("url-short", user.ID, time.Now().Add(time.Hour*24), "access")

	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.GenerateToken("url-short", user.ID, time.Now().Add(time.Hour*24), "refresh")

	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
