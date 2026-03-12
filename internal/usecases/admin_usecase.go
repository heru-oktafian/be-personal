package usecases

import (
	"context"

	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/heru-oktafian/be-personal/internal/repositories"
)

type AdminUseCase interface {
	GetProfile(ctx context.Context, id string) (*entities.Admin, error)
}

type adminUseCase struct {
	adminRepo repositories.AdminRepository
}

func NewAdminUseCase(adminRepo repositories.AdminRepository) AdminUseCase {
	return &adminUseCase{adminRepo: adminRepo}
}

func (u *adminUseCase) GetProfile(ctx context.Context, id string) (*entities.Admin, error) {
	return u.adminRepo.GetByID(ctx, id)
}
