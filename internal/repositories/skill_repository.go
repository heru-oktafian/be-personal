package repositories

import (
	"context"

	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/jmoiron/sqlx"
)

type SkillRepository interface {
	Create(ctx context.Context, skill *entities.Skill) error
	GetAll(ctx context.Context) ([]entities.Skill, error)
	GetByID(ctx context.Context, id string) (*entities.Skill, error)
	Update(ctx context.Context, skill *entities.Skill) error
	Delete(ctx context.Context, id string) error
}

type skillRepository struct {
	db *sqlx.DB
}

func NewSkillRepository(db *sqlx.DB) SkillRepository {
	return &skillRepository{db: db}
}

func (r *skillRepository) Create(ctx context.Context, skill *entities.Skill) error {
	query := `INSERT INTO skills (id, name, category, percentage, icon_url, order_num) VALUES (:id, :name, :category, :percentage, :icon_url, :order_num)`
	_, err := r.db.NamedExecContext(ctx, query, skill)
	return err
}

// ... Implementasi GetAll, GetByID, Update, dan Delete untuk Skill serupa dengan Service (hanya beda tabel dan field)
func (r *skillRepository) GetAll(ctx context.Context) ([]entities.Skill, error) {
	var skills []entities.Skill
	err := r.db.SelectContext(ctx, &skills, `SELECT id, name, category, percentage, COALESCE(icon_url, '') as icon_url, order_num FROM skills ORDER BY category, order_num ASC`)
	return skills, err
}

func (r *skillRepository) GetByID(ctx context.Context, id string) (*entities.Skill, error) {
	var skill entities.Skill
	err := r.db.GetContext(ctx, &skill, `SELECT id, name, category, percentage, COALESCE(icon_url, '') as icon_url, order_num FROM skills WHERE id = $1`, id)
	return &skill, err
}

func (r *skillRepository) Update(ctx context.Context, skill *entities.Skill) error {
	query := `UPDATE skills SET name = :name, category = :category, percentage = :percentage, icon_url = :icon_url, order_num = :order_num WHERE id = :id`
	_, err := r.db.NamedExecContext(ctx, query, skill)
	return err
}

func (r *skillRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM skills WHERE id = $1`, id)
	return err
}
