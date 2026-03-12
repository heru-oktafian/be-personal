package repositories

import (
	"context"

	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/jmoiron/sqlx"
)

type SeoRepository interface {
	Upsert(ctx context.Context, seo *entities.SeoMetadata) error
	GetByReference(ctx context.Context, refType, refID string) (*entities.SeoMetadata, error)
}

type seoRepository struct {
	db *sqlx.DB
}

func NewSeoRepository(db *sqlx.DB) SeoRepository {
	return &seoRepository{db: db}
}

func (r *seoRepository) Upsert(ctx context.Context, seo *entities.SeoMetadata) error {
	query := `
		INSERT INTO seo_metadata (id, reference_id, reference_type, meta_title, meta_desc, og_image_url, alt_text)
		VALUES (:id, :reference_id, :reference_type, :meta_title, :meta_desc, :og_image_url, :alt_text)
		ON CONFLICT (reference_id, reference_type)
		DO UPDATE SET
			meta_title = EXCLUDED.meta_title,
			meta_desc = EXCLUDED.meta_desc,
			og_image_url = EXCLUDED.og_image_url,
			alt_text = EXCLUDED.alt_text;
	`
	_, err := r.db.NamedExecContext(ctx, query, seo)
	return err
}

func (r *seoRepository) GetByReference(ctx context.Context, refType, refID string) (*entities.SeoMetadata, error) {
	var seo entities.SeoMetadata
	query := `SELECT * FROM seo_metadata WHERE reference_type = $1 AND reference_id = $2 LIMIT 1`
	err := r.db.GetContext(ctx, &seo, query, refType, refID)
	if err != nil {
		return nil, err
	}
	return &seo, nil
}
