package repositories

import (
	"context"

	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/jmoiron/sqlx"
)

type ProjectRepository interface {
	FetchAll(ctx context.Context) ([]entities.Project, error)
	GetBySlug(ctx context.Context, slug string) (*entities.Project, error)
	Create(ctx context.Context, project *entities.Project) error
	GetByID(ctx context.Context, id string) (*entities.Project, error)
	Update(ctx context.Context, project *entities.Project) error
	Delete(ctx context.Context, id string) error
}

type projectRepository struct {
	db *sqlx.DB
}

func NewProjectRepository(db *sqlx.DB) ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) FetchAll(ctx context.Context) ([]entities.Project, error) {
	var projects []entities.Project
	query := `SELECT id, title, slug, description, image_url, is_published, created_at, updated_at FROM projects ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &projects, query)
	return projects, err
}

func (r *projectRepository) GetBySlug(ctx context.Context, slug string) (*entities.Project, error) {
	var project entities.Project
	query := `SELECT id, title, slug, description, image_url, is_published, created_at, updated_at FROM projects WHERE slug = $1 LIMIT 1`
	err := r.db.GetContext(ctx, &project, query, slug)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *projectRepository) Create(ctx context.Context, project *entities.Project) error {
	query := `
		INSERT INTO projects (id, title, slug, description, image_url, is_published, created_at, updated_at)
		VALUES (:id, :title, :slug, :description, :image_url, :is_published, :created_at, :updated_at)
	`
	_, err := r.db.NamedExecContext(ctx, query, project)
	return err
}

func (r *projectRepository) GetByID(ctx context.Context, id string) (*entities.Project, error) {
	var project entities.Project
	query := `SELECT id, title, slug, description, image_url, is_published, created_at, updated_at FROM projects WHERE id = $1 LIMIT 1`
	err := r.db.GetContext(ctx, &project, query, id)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *projectRepository) Update(ctx context.Context, project *entities.Project) error {
	query := `
		UPDATE projects 
		SET title = :title, slug = :slug, description = :description, image_url = :image_url, 
		    is_published = :is_published, updated_at = :updated_at
		WHERE id = :id
	`
	_, err := r.db.NamedExecContext(ctx, query, project)
	return err
}

func (r *projectRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM projects WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
