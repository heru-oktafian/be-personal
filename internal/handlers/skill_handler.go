package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/heru-oktafian/be-personal/internal/usecases"
)

type SkillHandler struct {
	uc usecases.SkillUseCase
}

func NewSkillHandler(uc usecases.SkillUseCase) *SkillHandler {
	return &SkillHandler{uc: uc}
}

func (h *SkillHandler) GetAll(c *fiber.Ctx) error {
	res, err := h.uc.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": res})
}

func (h *SkillHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.uc.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Keahlian teknis tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"data": res})
}

func (h *SkillHandler) Create(c *fiber.Ctx) error {
	var payload entities.Skill
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format request tidak valid"})
	}

	if err := h.uc.Create(c.Context(), &payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Keahlian teknis berhasil ditambahkan",
		"data":    payload,
	})
}

func (h *SkillHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	// Pastikan data eksis terlebih dahulu sebelum di-update
	existingSkill, err := h.uc.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Keahlian teknis tidak ditemukan"})
	}

	if err := c.BodyParser(existingSkill); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format request tidak valid"})
	}

	existingSkill.ID = id // Proteksi agar ID tidak berubah dari payload berbahaya

	if err := h.uc.Update(c.Context(), existingSkill); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal memperbarui keahlian teknis"})
	}

	return c.JSON(fiber.Map{
		"message": "Keahlian teknis berhasil diperbarui",
		"data":    existingSkill,
	})
}

func (h *SkillHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.uc.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus keahlian teknis"})
	}

	return c.JSON(fiber.Map{"message": "Keahlian teknis berhasil dihapus"})
}
