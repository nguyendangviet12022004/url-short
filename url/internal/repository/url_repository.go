package repository

import (
	"context"
	"nguyendangviet12022004/url-short/url/internal/model"

	"gorm.io/gorm"
)

type UrlRepository interface {
	Create(ctx context.Context, url *model.Url) error
	FindByShortUrl(ctx context.Context, shortUrl string) (*model.Url, error)
	FindByUserId(ctx context.Context, userId uint) ([]model.Url, error)
	Delete(ctx context.Context, id uint) (int, error)
}

type urlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) UrlRepository {
	return &urlRepository{db: db}
}

func (r *urlRepository) Create(ctx context.Context, url *model.Url) error {
	return gorm.G[model.Url](r.db).Create(ctx, url)
}

func (r *urlRepository) FindByShortUrl(ctx context.Context, shortUrl string) (*model.Url, error) {
	url, err := gorm.G[model.Url](r.db).Where("short_url = ?", shortUrl).First(ctx)
	if err != nil {
		return nil, err
	}
	return &url, nil
}

func (r *urlRepository) FindByUserId(ctx context.Context, userId uint) ([]model.Url, error) {
	urls, err := gorm.G[model.Url](r.db).Where("user_id = ?", userId).Find(ctx)
	if err != nil {
		return nil, err
	}
	return urls, nil
}

func (r *urlRepository) Delete(ctx context.Context, id uint) (int, error) {
	return gorm.G[model.Url](r.db).Where("id = ?", id).Delete(ctx)
}
