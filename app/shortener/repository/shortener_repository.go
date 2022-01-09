package repository

import (
	"context"
	"database/sql"
	"errors"
	"yuuki/domain"
	"yuuki/pkg/helper"
)

type shortenerRepository struct {
	db *sql.DB
}

func NewShortenerRepository(db *sql.DB) *shortenerRepository {
	return &shortenerRepository{db: db}
}

func (repository *shortenerRepository) Create(ctx context.Context, shortener domain.Shortener) domain.Shortener {
	statement := "INSERT INTO shorteners(real_url, slug) VALUE (?, ?);"
	result, err := repository.db.ExecContext(ctx, statement, shortener.RealURL, shortener.Slug)
	helper.PanicIfErr(err)

	id, err := result.LastInsertId()
	helper.PanicIfErr(err)

	shortener.ID = int(id)
	return shortener
}

func (repository *shortenerRepository) Delete(ctx context.Context, id int) domain.Shortener {
	panic("implement me")
}

func (repository *shortenerRepository) FindBy(ctx context.Context, statement string, args []interface{}) (domain.Shortener, error) {
	rows, err := repository.db.QueryContext(ctx, statement, args...)
	helper.PanicIfErr(err)
	defer rows.Close()

	shortener := domain.Shortener{}
	if rows.Next() {
		helper.PanicIfErr(rows.Scan(&shortener.ID, &shortener.RealURL, &shortener.Slug))
		return domain.Shortener{
			ID:      shortener.ID,
			RealURL: shortener.RealURL,
			Slug:    shortener.Slug,
		}, nil
	} else {
		return shortener, errors.New("shortener is not found")
	}
}

func (repository *shortenerRepository) UpdateView(ctx context.Context, id int) {
	statement := "UPDATE shorteners SET views=views + 1, last_viewed=CURRENT_TIMESTAMP WHERE id = ?;"
	repository.db.ExecContext(ctx, statement, id)
}
