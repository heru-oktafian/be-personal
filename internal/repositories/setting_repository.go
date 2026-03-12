package repositories

import (
	"context"

	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/jmoiron/sqlx"
)

type SettingRepository interface {
	GetAll(ctx context.Context) ([]entities.SiteSetting, error)
	Upsert(ctx context.Context, setting *entities.SiteSetting) error
	GetByKey(ctx context.Context, key string) (*entities.SiteSetting, error)
	Delete(ctx context.Context, key string) error
}

type settingRepository struct {
	db *sqlx.DB
}

func NewSettingRepository(db *sqlx.DB) SettingRepository {
	return &settingRepository{db: db}
}

func (r *settingRepository) GetAll(ctx context.Context) ([]entities.SiteSetting, error) {
	var settings []entities.SiteSetting
	err := r.db.SelectContext(ctx, &settings, `SELECT key, value::text, updated_at FROM site_settings`)
	return settings, err
}

func (r *settingRepository) Upsert(ctx context.Context, setting *entities.SiteSetting) error {
	query := `
		INSERT INTO site_settings (key, value, updated_at) 
		VALUES (:key, :value, CURRENT_TIMESTAMP)
		ON CONFLICT (key) DO UPDATE 
		SET value = EXCLUDED.value, updated_at = CURRENT_TIMESTAMP
	`
	_, err := r.db.NamedExecContext(ctx, query, setting)
	return err
}

func (r *settingRepository) GetByKey(ctx context.Context, key string) (*entities.SiteSetting, error) {
	var setting entities.SiteSetting
	err := r.db.GetContext(ctx, &setting, `SELECT key, value::text, updated_at FROM site_settings WHERE key = $1`, key)
	return &setting, err
}

func (r *settingRepository) Delete(ctx context.Context, key string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM site_settings WHERE key = $1`, key)
	return err
}
