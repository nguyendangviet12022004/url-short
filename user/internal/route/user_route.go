package route

import (
	"nguyendangviet12022004/url-short/user/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitUserRoute(router *gin.Engine, userHandler handler.UserHandler) {
	userGroup := router.Group("/user")
	userGroup.POST("/register", userHandler.RegisterUser)
	userGroup.POST("/login", userHandler.LoginUser)
}
