package dto

type CreateUrlRequest struct {
	Title   string `json:"title" binding:"required"`
	LongUrl string `json:"long_url" binding:"required"`
}
