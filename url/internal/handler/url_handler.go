package handler

import (
	"nguyendangviet12022004/url-short/url/internal/dto"
	"nguyendangviet12022004/url-short/url/internal/exception"
	"nguyendangviet12022004/url-short/url/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UrlHandler interface {
	Create(ctx *gin.Context)
	GetFullUrl(ctx *gin.Context)
	GetUrlsByUserId(ctx *gin.Context)
}

type urlHandler struct {
	urlService service.UrlService
}

func NewUrlHandler(urlService service.UrlService) UrlHandler {
	return &urlHandler{urlService: urlService}
}

func (h *urlHandler) Create(ctx *gin.Context) {
	var req dto.CreateUrlRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		Error(ctx, err)
		return
	}

	userId := ctx.GetHeader("X-User-Id")
	userIdUint, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		Error(ctx, err)
		return
	}

	if err := h.urlService.Create(ctx, &req, uint(userIdUint)); err != nil {
		Error(ctx, err)
		return
	}

	Ok(ctx, gin.H{"message": "URL created successfully"})
}

func (h *urlHandler) GetFullUrl(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	if shortUrl == "" {
		Error(ctx, exception.ErrBadRequest)
		return
	}

	fullUrl, err := h.urlService.GetFullUrl(ctx, shortUrl)
	if err != nil {
		Error(ctx, err)
		return
	}

	Redirect(ctx, fullUrl)
}

func (h *urlHandler) GetUrlsByUserId(ctx *gin.Context) {
	userId := ctx.GetHeader("X-User-Id")
	userIdUint, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		Error(ctx, err)
		return
	}

	urls, err := h.urlService.GetUrlsByUserId(ctx, uint(userIdUint))
	if err != nil {
		Error(ctx, err)
		return
	}

	Ok(ctx, urls)
}
