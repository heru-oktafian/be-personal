package repositories

import (
	"context"

	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/jmoiron/sqlx"
)

type HomeRepository interface {
	GetAggregatedData(ctx context.Context) (*entities.HomeAggregatedData, error)
}

type homeRepository struct {
	db *sqlx.DB
}

func NewHomeRepository(db *sqlx.DB) HomeRepository {
	return &homeRepository{db: db}
}

func (r *homeRepository) GetAggregatedData(ctx context.Context) (*entities.HomeAggregatedData, error) {
	var data entities.HomeAggregatedData

	// 1. Ambil Settings (Hero, About, CTA, Contacts, Footer)
	err := r.db.SelectContext(ctx, &data.Settings, `SELECT key, value::text, updated_at FROM site_settings`)
	if err != nil {
		return nil, err
	}

	// 2. Ambil Services
	err = r.db.SelectContext(ctx, &data.Services, `SELECT id, title, description, icon_name, order_num FROM services ORDER BY order_num ASC`)
	if err != nil {
		return nil, err
	}

	// 3. Ambil Skills
	err = r.db.SelectContext(ctx, &data.Skills, `SELECT id, name, category, percentage, COALESCE(icon_url, '') as icon_url, order_num FROM skills ORDER BY category, order_num ASC`)
	if err != nil {
		return nil, err
	}

	// 4. Ambil 3 Project Teratas (Featured) untuk bagian "Portfolio Saya"
	err = r.db.SelectContext(ctx, &data.Projects, `SELECT id, title, slug, description, image_url, COALESCE(tags, '') as tags, is_published, created_at, updated_at FROM projects WHERE is_published = true ORDER BY created_at DESC LIMIT 3`)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
