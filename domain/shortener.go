package domain

import (
	"context"
	"time"
)

type Shortener struct {
	ID         int
	RealURL    string
	Slug       string
	Views      int
	LastViewed time.Time
	CreatedAt  time.Time
}

type ShortenerPayload struct {
	ID      int    `json:"id"`
	RealURL string `json:"real_url" validate:"required,url"`
	Slug    string `json:"slug" validate:"required"`
}

func (s ShortenerPayload) FillForNewRecord() Shortener {
	return Shortener{
		RealURL: s.RealURL,
		Slug:    s.Slug,
	}
}

func (s Shortener) AsPayload() ShortenerPayload {
	return ShortenerPayload{
		ID:      s.ID,
		RealURL: s.RealURL,
		Slug:    BaseUrl + "/api/shorteners/" + s.Slug,
	}
}

type ShortenerRepository interface {
	Create(ctx context.Context, shortener Shortener) Shortener
	Delete(ctx context.Context, id int) Shortener
	FindBy(ctx context.Context, statement string, args []interface{}) (Shortener, error)
	UpdateView(ctx context.Context, id int)
}

type ShortenerUsecase interface {
	Create(ctx context.Context, payload ShortenerPayload) ShortenerPayload
	GetBySlug(ctx context.Context, payload ShortenerPayload) ShortenerPayload
}
