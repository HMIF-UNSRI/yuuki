package domain

import (
	"context"
	"time"
)

type Category struct {
	ID        int
	Name      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CategoryPayload struct {
	ID   int
	Name string `json:"name" validate:"required"`
	Slug string `json:"slug"`
}

func (c CategoryPayload) FillForNewRecord() Category {
	return Category{
		Name: c.Name,
		Slug: c.Slug,
	}
}

func (c Category) AsPayload() CategoryPayload {
	return CategoryPayload{
		ID:   c.ID,
		Name: c.Name,
		Slug: c.Slug,
	}
}

type CategoryRepository interface {
	Create(ctx context.Context, category Category) Category
	Update(ctx context.Context, id int, category Category) Category
	Delete(ctx context.Context, id int) Category
	FindBy(ctx context.Context, statement string, args []interface{}) (Category, error)
	FindAll(ctx context.Context, statement string) []Category
}

type CategoryUsecase interface {
	Create(ctx context.Context, payload CategoryPayload) CategoryPayload
}
