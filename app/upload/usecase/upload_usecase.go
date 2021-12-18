package usecase

import (
	"context"
	"yuuki/domain"
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
