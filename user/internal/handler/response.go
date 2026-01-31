package handler

import (
	"errors"
	"net/http"
	"nguyendangviet12022004/url-short/user/internal/exception"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

func Error(c *gin.Context, err error) {

	var customErr *exception.CustomError
	if !errors.As(err, &customErr) {
		customErr = exception.ErrInternalServer
	}

	c.JSON(customErr.Code, gin.H{
		"error_code": customErr.ErrorCode,
		"message":    customErr.Message,
	})
}
