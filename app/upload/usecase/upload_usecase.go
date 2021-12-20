package usecase

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"yuuki/domain"
	"yuuki/pkg/helper"
)

type uploadUsecase struct {
	uploadRepository domain.UploadRepository
}

func NewUploadUsecase(uploadRepository domain.UploadRepository) *uploadUsecase {
	return &uploadUsecase{uploadRepository: uploadRepository}
}

func (usecase *uploadUsecase) Create(ctx context.Context, payload domain.UploadPayload) domain.UploadPayload {
	upload := usecase.uploadRepository.Create(ctx, payload.FillForNewRecord())
	return upload.AsPayload()
}

func (usecase *uploadUsecase) List(ctx context.Context, param domain.PaginationParam) ([]domain.UploadPayload, domain.Pagination) {
	queryBuilder := sq.Select("id", "image_name", "alt_text").From("uploads").
		OrderBy("id DESC")

	if param.Limit > 0 {
		queryBuilder = queryBuilder.Limit(uint64(param.Limit))
	} else {
		queryBuilder = queryBuilder.Limit(domain.DefaultPaginationLimit)
	}

	if param.CursorID == 1 {
		panic(domain.NewBadRequestError("cursor cannot be 1"))
	} else if param.CursorID > 1 {
		queryBuilder = queryBuilder.Where(sq.Lt{
			"id": param.CursorID,
		})
	}
	statement, args, err := queryBuilder.ToSql()
	helper.PanicIfErr(err)

	uploads := usecase.uploadRepository.FindAll(ctx, statement, args)

	var uploadResponses []domain.UploadPayload
	for _, upload := range uploads {
		uploadResponses = append(uploadResponses, upload.AsPayload())
	}

	count := len(uploadResponses)
	pagination := domain.Pagination{
		Count: uint32(count),
		Next: domain.BaseUrl + fmt.Sprintf("/api/uploads?limit=%d&cursor=%d",
			domain.DefaultPaginationLimit, uploadResponses[count-1].ID),
	}

	if param.CursorID != 0 && param.Limit != 0 {
		pagination.Previous = domain.BaseUrl + fmt.Sprintf("/api/uploads?limit=%d&cursor=%d",
			param.Limit, uploadResponses[0].ID+int(param.Limit)+1)
	}

	return uploadResponses, pagination
}
