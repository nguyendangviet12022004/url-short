package handler

import (
	"errors"
	"net/http"
	"nguyendangviet12022004/url-short/url/internal/exception"

	"github.com/gin-gonic/gin"
)

func Ok(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "OK",
		"data":    data,
	})
}

func Error(ctx *gin.Context, err error) {
	var customError *exception.CustomError
	if errors.As(err, &customError) {
		ctx.JSON(customError.HttpStatus, gin.H{
			"code":    customError.Code,
			"message": customError.Message,
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code":    500,
		"message": err.Error(),
	})
}

func Redirect(ctx *gin.Context, url string) {
	ctx.Redirect(http.StatusSeeOther, url)
}
