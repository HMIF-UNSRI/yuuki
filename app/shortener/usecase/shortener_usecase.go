package usecase

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-playground/validator"
	"yuuki/domain"
	"yuuki/pkg/helper"
)

type shortenerUsecase struct {
	shortenerRepository domain.ShortenerRepository
	validate            *validator.Validate
}

func NewShortenerUsecase(shortenerRepository domain.ShortenerRepository, validate *validator.Validate) *shortenerUsecase {
	return &shortenerUsecase{shortenerRepository: shortenerRepository, validate: validate}
}

func (usecase *shortenerUsecase) Create(ctx context.Context, payload domain.ShortenerPayload) domain.ShortenerPayload {
	helper.PanicIfErr(usecase.validate.Struct(&payload))

	statement, args, err := sq.Select("id", "real_url", "slug").From("shorteners").
		Limit(1).Where(sq.Eq{"slug": payload.Slug}).ToSql()
	helper.PanicIfErr(err)

	category, err := usecase.shortenerRepository.FindBy(ctx, statement, args)
	if err == nil {
		panic(domain.NewAlreadyExistError("shortener is already exist"))
	}

	category = usecase.shortenerRepository.Create(ctx, payload.FillForNewRecord())
	return category.AsPayload()
}

func (usecase *shortenerUsecase) GetBySlug(ctx context.Context, payload domain.ShortenerPayload) domain.ShortenerPayload {
	statement, args, err := sq.Select("id", "real_url", "slug").From("shorteners").
		Limit(1).Where(sq.Eq{"slug": payload.Slug}).ToSql()
	helper.PanicIfErr(err)

	category, err := usecase.shortenerRepository.FindBy(ctx, statement, args)
	if err != nil {
		panic(domain.NewNotFoundError("shortener is not found"))
	} else {
		// Asynchronous
		go usecase.shortenerRepository.UpdateView(ctx, category.ID)

		return category.AsPayload()
	}
}
