package service

import (
	"context"
	"nguyendangviet12022004/url-short/url/internal/dto"
	"nguyendangviet12022004/url-short/url/internal/exception"
	"nguyendangviet12022004/url-short/url/internal/model"
	"nguyendangviet12022004/url-short/url/internal/repository"

	"github.com/google/uuid"
)

type UrlService interface {
	Create(ctx context.Context, request *dto.CreateUrlRequest, userId uint) error
	GetFullUrl(ctx context.Context, shortUrl string) (string, error)
	GetUrlsByUserId(ctx context.Context, userId uint) ([]model.Url, error)
}

type urlService struct {
	urlRepository repository.UrlRepository
}

func NewUrlService(urlRepository repository.UrlRepository) UrlService {
	return &urlService{urlRepository: urlRepository}
}

func (s *urlService) Create(ctx context.Context, req *dto.CreateUrlRequest, userId uint) error {

	shortUrl := generateShortUrl()

	for {
		existShortUrl, _ := s.urlRepository.FindByShortUrl(ctx, shortUrl)
		if existShortUrl != nil {
			shortUrl = generateShortUrl()
			continue
		}
		break
	}

	url := &model.Url{
		Title:    req.Title,
		LongUrl:  req.LongUrl,
		UserId:   userId,
		ShortUrl: shortUrl,
	}

	return s.urlRepository.Create(ctx, url)
}

func generateShortUrl() string {
	return uuid.New().String()
}

func (s *urlService) GetFullUrl(ctx context.Context, shortUrl string) (string, error) {
	url, err := s.urlRepository.FindByShortUrl(ctx, shortUrl)
	if err != nil {
		return "", exception.ErrShortUrlNotFound
	}
	return url.LongUrl, nil
}

func (s *urlService) GetUrlsByUserId(ctx context.Context, userId uint) ([]model.Url, error) {
	return s.urlRepository.FindByUserId(ctx, userId)
}
