package repository

import (
	"context"
	"database/sql"
	"errors"
	"yuuki/domain"
	"yuuki/pkg/helper"
)

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (repository *categoryRepository) Create(ctx context.Context, category domain.Category) domain.Category {
	statement := "INSERT INTO categories(name, slug) VALUE (?, ?);"
	result, err := repository.db.ExecContext(ctx, statement, category.Name, category.Slug)
	helper.PanicIfErr(err)

	id, err := result.LastInsertId()
	helper.PanicIfErr(err)

	category.ID = int(id)
	return category
}

func (repository *categoryRepository) Update(ctx context.Context, category domain.Category) {
	statement := "UPDATE categories SET name=?, slug=?, updated_at=CURRENT_TIMESTAMP WHERE id = ?;"
	_, err := repository.db.ExecContext(ctx, statement, category.Name, category.Slug, category.ID)
	helper.PanicIfErr(err)
}

func (repository *categoryRepository) Delete(ctx context.Context, id int) domain.Category {
	panic("implement me")
}

func (repository *categoryRepository) FindBy(ctx context.Context, statement string, args []interface{}) (domain.Category, error) {
	rows, err := repository.db.QueryContext(ctx, statement, args...)
	helper.PanicIfErr(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		helper.PanicIfErr(rows.Scan(&category.ID, &category.Name, &category.Slug))
		return domain.Category{
			ID:   category.ID,
			Name: category.Name,
			Slug: category.Slug,
		}, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *categoryRepository) FindAll(ctx context.Context) []domain.Category {
	statement := "SELECT id, name, slug FROM categories;"
	rows, err := repository.db.QueryContext(ctx, statement)
	helper.PanicIfErr(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		helper.PanicIfErr(rows.Scan(&category.ID, &category.Name, &category.Slug))
		categories = append(categories, category)
	}

	return categories
}
