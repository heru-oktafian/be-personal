package usecases

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/heru-oktafian/be-personal/internal/repositories"
)

type ContactUseCase interface {
	SubmitMessage(ctx context.Context, contact *entities.Contact) error
	GetAllMessages(ctx context.Context) ([]entities.Contact, error)
	GetMessageByID(ctx context.Context, id string) (*entities.Contact, error)
	DeleteMessage(ctx context.Context, id string) error
}

type contactUseCase struct {
	contactRepo repositories.ContactRepository
}

func NewContactUseCase(contactRepo repositories.ContactRepository) ContactUseCase {
	return &contactUseCase{contactRepo: contactRepo}
}

func (u *contactUseCase) SubmitMessage(ctx context.Context, contact *entities.Contact) error {
	contact.ID = uuid.New().String()
	contact.CreatedAt = time.Now()
	return u.contactRepo.Create(ctx, contact)
}

func (u *contactUseCase) GetAllMessages(ctx context.Context) ([]entities.Contact, error) {
	return u.contactRepo.FetchAll(ctx)
}

func (u *contactUseCase) GetMessageByID(ctx context.Context, id string) (*entities.Contact, error) {
	return u.contactRepo.GetByID(ctx, id)
}

func (u *contactUseCase) DeleteMessage(ctx context.Context, id string) error {
	return u.contactRepo.Delete(ctx, id)
}
