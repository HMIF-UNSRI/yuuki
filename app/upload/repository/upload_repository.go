package repository

import (
	"context"
	"database/sql"
	"yuuki/domain"
	"yuuki/pkg/helper"
)

type uploadRepository struct {
	db *sql.DB
}

func NewUploadRepository(db *sql.DB) *uploadRepository {
	return &uploadRepository{db: db}
}

func (repository *uploadRepository) Create(ctx context.Context, upload domain.Upload) domain.Upload {
	statement := "INSERT INTO uploads(image_name, alt_text) VALUE (?, ?);"
	result, err := repository.db.ExecContext(ctx, statement, upload.ImageName, upload.AltText)
	helper.PanicIfErr(err)

	id, err := result.LastInsertId()
	helper.PanicIfErr(err)

	upload.ID = int(id)
	return upload
}
