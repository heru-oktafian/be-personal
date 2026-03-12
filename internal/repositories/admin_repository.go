package repositories

import (
	"context"

	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/jmoiron/sqlx"
)

type AdminRepository interface {
	GetByEmail(ctx context.Context, email string) (*entities.Admin, error)
	GetByID(ctx context.Context, id string) (*entities.Admin, error)
}

type adminRepository struct {
	db *sqlx.DB
}

func NewAdminRepository(db *sqlx.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) GetByEmail(ctx context.Context, email string) (*entities.Admin, error) {
	var admin entities.Admin
	query := `SELECT id, email, password_hash, created_at FROM admins WHERE email = $1 LIMIT 1`
	err := r.db.GetContext(ctx, &admin, query, email)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

// GetByID(ctx context.Context, id string) (*entities.Admin, error)
func (r *adminRepository) GetByID(ctx context.Context, id string) (*entities.Admin, error) {
	var admin entities.Admin
	query := `SELECT id, email, created_at FROM admins WHERE id = $1 LIMIT 1`
	err := r.db.GetContext(ctx, &admin, query, id)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
