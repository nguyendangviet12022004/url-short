package handler

import (
	"context"
	"nguyendangviet12022004/url-short/user/internal/dto"
	"nguyendangviet12022004/url-short/user/internal/exception"
	"nguyendangviet12022004/url-short/user/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService: userService}
}

func (uh *userHandler) RegisterUser(c *gin.Context) {

	var request *dto.RegisterRequest
	if err := c.BindJSON(&request); err != nil {
		Error(c, exception.ErrInvalidRequestBody)
		return
	}

	ctx := context.Background()
	err := uh.userService.RegisterUser(ctx, request)
	if err != nil {
		Error(c, err)
	} else {
		Success(c, gin.H{
			"message": "User registered successfully",
		})
	}
}

func (uh *userHandler) LoginUser(c *gin.Context) {

	var request *dto.LoginRequest
	if err := c.BindJSON(&request); err != nil {
		Error(c, exception.ErrInvalidRequestBody)
		return
	}

	ctx := context.Background()
	loginResponse, err := uh.userService.Login(ctx, request)

	if err != nil {
		Error(c, err)
	} else {
		c.SetCookie("access_token", loginResponse.AccessToken, 60*60, "/", "", false, true)
		c.SetCookie("refresh_token", loginResponse.RefreshToken, 60*60, "/", "", false, true)
		Success(c, gin.H{
			"message": "User logged in successfully",
		})
	}
}
