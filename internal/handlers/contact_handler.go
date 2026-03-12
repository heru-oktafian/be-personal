package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/heru-oktafian/be-personal/internal/usecases"
)

type ContactHandler struct {
	contactUseCase usecases.ContactUseCase
}

func NewContactHandler(contactUseCase usecases.ContactUseCase) *ContactHandler {
	return &ContactHandler{contactUseCase: contactUseCase}
}

func (h *ContactHandler) Submit(c *fiber.Ctx) error {
	var contact entities.Contact
	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid"})
	}

	if err := h.contactUseCase.SubmitMessage(c.Context(), &contact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal mengirim pesan"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Pesan berhasil dikirim"})
}

func (h *ContactHandler) GetAll(c *fiber.Ctx) error {
	contacts, err := h.contactUseCase.GetAllMessages(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal mengambil pesan"})
	}
	return c.JSON(fiber.Map{"data": contacts})
}

func (h *ContactHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	contact, err := h.contactUseCase.GetMessageByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pesan tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"data": contact})
}

func (h *ContactHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.contactUseCase.DeleteMessage(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus pesan"})
	}
	return c.JSON(fiber.Map{"message": "Pesan berhasil dihapus"})
}
