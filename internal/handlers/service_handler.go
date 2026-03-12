package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/heru-oktafian/be-personal/internal/usecases"
)

// --- SERVICE HANDLER ---
type ServiceHandler struct {
	uc usecases.ServiceUseCase
}

func NewServiceHandler(uc usecases.ServiceUseCase) *ServiceHandler {
	return &ServiceHandler{uc: uc}
}

func (h *ServiceHandler) GetAll(c *fiber.Ctx) error {
	res, err := h.uc.GetAll(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": res})
}

func (h *ServiceHandler) Create(c *fiber.Ctx) error {
	var payload entities.Service
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format request tidak valid"})
	}
	if err := h.uc.Create(c.Context(), &payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal membuat layanan"})
	}
	return c.JSON(fiber.Map{"message": "Layanan berhasil ditambahkan", "data": payload})
}

func (h *ServiceHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.uc.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Layanan tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"data": res})
}

func (h *ServiceHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	// Pastikan data eksis terlebih dahulu sebelum di-update
	existingService, err := h.uc.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Layanan tidak ditemukan"})
	}

	if err := c.BodyParser(existingService); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format request tidak valid"})
	}

	existingService.ID = id // Proteksi agar ID tidak berubah dari payload berbahaya

	if err := h.uc.Update(c.Context(), existingService); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal memperbarui layanan"})
	}

	return c.JSON(fiber.Map{
		"message": "Layanan berhasil diperbarui",
		"data":    existingService,
	})
}

func (h *ServiceHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.uc.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus layanan"})
	}

	return c.JSON(fiber.Map{"message": "Layanan berhasil dihapus"})
}
