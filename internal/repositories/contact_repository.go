package repositories

import (
	"context"

	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/jmoiron/sqlx"
)

type ContactRepository interface {
	Create(ctx context.Context, contact *entities.Contact) error
	FetchAll(ctx context.Context) ([]entities.Contact, error)
	GetByID(ctx context.Context, id string) (*entities.Contact, error)
	Delete(ctx context.Context, id string) error
}

type contactRepository struct {
	db *sqlx.DB
}

func NewContactRepository(db *sqlx.DB) ContactRepository {
	return &contactRepository{db: db}
}

func (r *contactRepository) Create(ctx context.Context, contact *entities.Contact) error {
	query := `
		INSERT INTO contacts (id, name, email, message, created_at)
		VALUES (:id, :name, :email, :message, :created_at)
	`
	_, err := r.db.NamedExecContext(ctx, query, contact)
	return err
}

func (r *contactRepository) FetchAll(ctx context.Context) ([]entities.Contact, error) {
	var contacts []entities.Contact
	query := `SELECT id, name, email, message, created_at FROM contacts ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &contacts, query)
	return contacts, err
}

func (r *contactRepository) GetByID(ctx context.Context, id string) (*entities.Contact, error) {
	var contact entities.Contact
	query := `SELECT id, name, email, message, created_at FROM contacts WHERE id = $1 LIMIT 1`
	err := r.db.GetContext(ctx, &contact, query, id)
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (r *contactRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM contacts WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
