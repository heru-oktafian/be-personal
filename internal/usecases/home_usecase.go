package usecases

import (
	"context"

	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/heru-oktafian/be-personal/internal/repositories"
)

type HomeUseCase interface {
	GetHomeData(ctx context.Context) (*entities.HomeAggregatedData, error)
}

type homeUseCase struct {
	homeRepo repositories.HomeRepository
}

func NewHomeUseCase(homeRepo repositories.HomeRepository) HomeUseCase {
	return &homeUseCase{homeRepo: homeRepo}
}

func (u *homeUseCase) GetHomeData(ctx context.Context) (*entities.HomeAggregatedData, error) {
	return u.homeRepo.GetAggregatedData(ctx)
}
