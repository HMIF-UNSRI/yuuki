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

	if payload.Slug == "" {
		slug := helper.ConvertNameToSlug(payload.Name)
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

func (usecase *categoryUsecase) Update(ctx context.Context, payload domain.CategoryPayload) domain.CategoryPayload {
	helper.PanicIfErr(usecase.validate.Struct(&payload))

	// Check slug
	if payload.Slug == "" {
		slug := helper.ConvertNameToSlug(payload.Name)
		payload.Slug = slug
	}

	statement, args, err := sq.Select("id", "name", "slug").From("categories").
		Limit(1).Where(sq.Eq{"slug": payload.Slug}).ToSql()
	helper.PanicIfErr(err)

	_, err = usecase.categoryRepository.FindBy(ctx, statement, args)
	if err == nil {
		panic(domain.NewAlreadyExistError("category is already exist"))
	}

	// Check category
	statement, args, err = sq.Select("id", "name", "slug").From("categories").
		Limit(1).Where(sq.Eq{"id": payload.ID}).ToSql()
	helper.PanicIfErr(err)

	_, err = usecase.categoryRepository.FindBy(ctx, statement, args)
	if err != nil {
		panic(domain.NewNotFoundError("category is not found"))
	}

	usecase.categoryRepository.Update(ctx, payload.FillForNewRecord())
	return payload
}

func (usecase *categoryUsecase) GetBy(ctx context.Context, payload domain.CategoryPayload) domain.CategoryPayload {
	var builder sq.SelectBuilder
	if payload.Slug != "" {
		builder = sq.Select("id", "name", "slug").From("categories").
			Limit(1).Where(sq.Eq{"slug": payload.Slug})
	} else if payload.ID != 0 {
		builder = sq.Select("id", "name", "slug").From("categories").
			Limit(1).Where(sq.Eq{"id": payload.ID})
	}

	statement, args, err := builder.ToSql()
	helper.PanicIfErr(err)

	category, err := usecase.categoryRepository.FindBy(ctx, statement, args)
	if err != nil {
		panic(domain.NewNotFoundError("category is not found"))
	}

	return category.AsPayload()
}

func (usecase *categoryUsecase) List(ctx context.Context) []domain.CategoryPayload {
	categories := usecase.categoryRepository.FindAll(ctx)
	var categoryResponses []domain.CategoryPayload
	for _, category := range categories {
		categoryResponses = append(categoryResponses, category.AsPayload())
	}
	return categoryResponses
}
