package domain

import (
	"context"
	"time"
)

const (
	BaseUrl      = "http://localhost:8080"
	MaxImageSize = 1024 * 1024 * 2 // 2MB
)

type Upload struct {
	ID        int
	ImageName string
	AltText   string
	CreatedAt time.Time
}

type UploadPayload struct {
	ID        int    `json:"id"`
	ImageName string `json:"image_name"`
	ImageURL  string `json:"image_url"`
	AltText   string `json:"alt_text"`
}

func (u UploadPayload) FillForNewRecord() Upload {
	return Upload{
		ImageName: u.ImageName,
		AltText:   u.AltText,
	}
}

func (u Upload) AsPayload() UploadPayload {
	return UploadPayload{
		ID:        u.ID,
		ImageName: u.ImageName,
		ImageURL:  BaseUrl + "/api/resources/" + u.ImageName,
		AltText:   u.AltText,
	}
}

type UploadRepository interface {
	Create(ctx context.Context, upload Upload) Upload
	FindAll(ctx context.Context, statement string, args []interface{}) []Upload
}

type UploadUsecase interface {
	Create(ctx context.Context, payload UploadPayload) UploadPayload
	List(ctx context.Context, param PaginationParam) ([]UploadPayload, Pagination)
}
