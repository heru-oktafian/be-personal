package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/heru-oktafian/be-personal/internal/repositories"
)

type SeoUseCase interface {
	SaveSeoMetadata(ctx context.Context, seo *entities.SeoMetadata) error
	GetSeoByReference(ctx context.Context, refType, refID string) (*entities.SeoMetadata, error)
}

type seoUseCase struct {
	seoRepo repositories.SeoRepository
}

func NewSeoUseCase(seoRepo repositories.SeoRepository) SeoUseCase {
	return &seoUseCase{seoRepo: seoRepo}
}

func (u *seoUseCase) SaveSeoMetadata(ctx context.Context, seo *entities.SeoMetadata) error {
	if seo.ID == "" {
		seo.ID = uuid.New().String()
	}
	return u.seoRepo.Upsert(ctx, seo)
}

func (u *seoUseCase) GetSeoByReference(ctx context.Context, refType, refID string) (*entities.SeoMetadata, error) {
	return u.seoRepo.GetByReference(ctx, refType, refID)
}
