package usecase

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-playground/validator"
	"yuuki/domain"
	"yuuki/pkg/helper"
)

type categoryUsecase struct {
	categoryRepository domain.CategoryRepository
	validate           *validator.Validate
}

func NewCategoryUsecase(categoryRepository domain.CategoryRepository, validate *validator.Validate) *categoryUsecase {
	return &categoryUsecase{categoryRepository: categoryRepository, validate: validate}
}

func (usecase *categoryUsecase) Create(ctx context.Context, payload domain.CategoryPayload) domain.CategoryPayload {
	helper.PanicIfErr(usecase.validate.Struct(&payload))

	var slug string
	if payload.Slug == "" {
		slug = helper.ConvertNameToSlug(payload.Name)
		payload.Slug = slug
	}

	statement, args, err := sq.Select("id", "name", "slug").From("categories").
		Limit(1).Where(sq.Eq{"slug": payload.Slug}).ToSql()
	helper.PanicIfErr(err)

	category, err := usecase.categoryRepository.FindBy(ctx, statement, args)
	if err == nil {
		panic(domain.NewAlreadyExistError("category is already exist"))
	}

	category = usecase.categoryRepository.Create(ctx, payload.FillForNewRecord())
	return category.AsPayload()
}
