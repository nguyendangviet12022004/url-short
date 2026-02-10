package route

import (
	"nguyendangviet12022004/url-short/url/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitUrlRoute(router *gin.Engine, urlHandler handler.UrlHandler) {
	urlGroup := router.Group("")
	urlGroup.POST("", urlHandler.Create)
	urlGroup.GET("/:shortUrl", urlHandler.GetFullUrl)
	urlGroup.GET("/", urlHandler.GetUrlsByUserId)
}
