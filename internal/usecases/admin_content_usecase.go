package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/heru-oktafian/be-personal/internal/repositories"
)

// --- SITE SETTINGS ---
type SettingUseCase interface {
	GetSettings(ctx context.Context) ([]entities.SiteSetting, error)
	SaveSetting(ctx context.Context, setting *entities.SiteSetting) error
	GetSettingByKey(ctx context.Context, key string) (*entities.SiteSetting, error)
	DeleteSetting(ctx context.Context, key string) error
}

type settingUseCase struct {
	repo repositories.SettingRepository
}

func NewSettingUseCase(repo repositories.SettingRepository) SettingUseCase {
	return &settingUseCase{repo}
}
func (u *settingUseCase) GetSettings(ctx context.Context) ([]entities.SiteSetting, error) {
	return u.repo.GetAll(ctx)
}
func (u *settingUseCase) SaveSetting(ctx context.Context, setting *entities.SiteSetting) error {
	return u.repo.Upsert(ctx, setting)
}

// --- SERVICES ---
type ServiceUseCase interface {
	GetAll(ctx context.Context) ([]entities.Service, error)
	GetByID(ctx context.Context, id string) (*entities.Service, error)
	Create(ctx context.Context, service *entities.Service) error
	Update(ctx context.Context, service *entities.Service) error
	Delete(ctx context.Context, id string) error
}

type serviceUseCase struct {
	repo repositories.ServiceRepository
}

func NewServiceUseCase(repo repositories.ServiceRepository) ServiceUseCase {
	return &serviceUseCase{repo}
}
func (u *serviceUseCase) GetAll(ctx context.Context) ([]entities.Service, error) {
	return u.repo.GetAll(ctx)
}
func (u *serviceUseCase) GetByID(ctx context.Context, id string) (*entities.Service, error) {
	return u.repo.GetByID(ctx, id)
}
func (u *serviceUseCase) Create(ctx context.Context, service *entities.Service) error {
	service.ID = uuid.New().String()
	return u.repo.Create(ctx, service)
}
func (u *serviceUseCase) Update(ctx context.Context, service *entities.Service) error {
	return u.repo.Update(ctx, service)
}
func (u *serviceUseCase) Delete(ctx context.Context, id string) error { return u.repo.Delete(ctx, id) }

// --- SKILLS ---
type SkillUseCase interface {
	GetAll(ctx context.Context) ([]entities.Skill, error)
	GetByID(ctx context.Context, id string) (*entities.Skill, error)
	Create(ctx context.Context, skill *entities.Skill) error
	Update(ctx context.Context, skill *entities.Skill) error
	Delete(ctx context.Context, id string) error
}

type skillUseCase struct{ repo repositories.SkillRepository }

func NewSkillUseCase(repo repositories.SkillRepository) SkillUseCase { return &skillUseCase{repo} }
func (u *skillUseCase) GetAll(ctx context.Context) ([]entities.Skill, error) {
	return u.repo.GetAll(ctx)
}
func (u *skillUseCase) GetByID(ctx context.Context, id string) (*entities.Skill, error) {
	return u.repo.GetByID(ctx, id)
}
func (u *skillUseCase) Create(ctx context.Context, skill *entities.Skill) error {
	skill.ID = uuid.New().String()
	return u.repo.Create(ctx, skill)
}
func (u *skillUseCase) Update(ctx context.Context, skill *entities.Skill) error {
	return u.repo.Update(ctx, skill)
}
func (u *skillUseCase) Delete(ctx context.Context, id string) error { return u.repo.Delete(ctx, id) }

// Tambahkan implementasinya
func (u *settingUseCase) GetSettingByKey(ctx context.Context, key string) (*entities.SiteSetting, error) {
	return u.repo.GetByKey(ctx, key)
}
func (u *settingUseCase) DeleteSetting(ctx context.Context, key string) error {
	return u.repo.Delete(ctx, key)
}
