package repositories

import (
	"context"

	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/jmoiron/sqlx"
)

type ServiceRepository interface {
	Create(ctx context.Context, service *entities.Service) error
	GetAll(ctx context.Context) ([]entities.Service, error)
	GetByID(ctx context.Context, id string) (*entities.Service, error)
	Update(ctx context.Context, service *entities.Service) error
	Delete(ctx context.Context, id string) error
}

type serviceRepository struct {
	db *sqlx.DB
}

func NewServiceRepository(db *sqlx.DB) ServiceRepository {
	return &serviceRepository{db: db}
}

func (r *serviceRepository) Create(ctx context.Context, service *entities.Service) error {
	query := `INSERT INTO services (id, title, description, icon_name, order_num) VALUES (:id, :title, :description, :icon_name, :order_num)`
	_, err := r.db.NamedExecContext(ctx, query, service)
	return err
}

func (r *serviceRepository) GetAll(ctx context.Context) ([]entities.Service, error) {
	var services []entities.Service
	err := r.db.SelectContext(ctx, &services, `SELECT id, title, description, icon_name, order_num FROM services ORDER BY order_num ASC`)
	return services, err
}

func (r *serviceRepository) GetByID(ctx context.Context, id string) (*entities.Service, error) {
	var service entities.Service
	err := r.db.GetContext(ctx, &service, `SELECT id, title, description, icon_name, order_num FROM services WHERE id = $1`, id)
	return &service, err
}

func (r *serviceRepository) Update(ctx context.Context, service *entities.Service) error {
	query := `UPDATE services SET title = :title, description = :description, icon_name = :icon_name, order_num = :order_num WHERE id = :id`
	_, err := r.db.NamedExecContext(ctx, query, service)
	return err
}

func (r *serviceRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM services WHERE id = $1`, id)
	return err
}
