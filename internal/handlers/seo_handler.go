package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/heru-oktafian/be-personal/internal/usecases"
)

type SeoHandler struct {
	seoUseCase usecases.SeoUseCase
}

func NewSeoHandler(seoUseCase usecases.SeoUseCase) *SeoHandler {
	return &SeoHandler{seoUseCase: seoUseCase}
}

func (h *SeoHandler) Upsert(c *fiber.Ctx) error {
	var seo entities.SeoMetadata
	if err := c.BodyParser(&seo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid"})
	}

	if err := h.seoUseCase.SaveSeoMetadata(c.Context(), &seo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan SEO Metadata"})
	}

	return c.JSON(fiber.Map{"message": "SEO Metadata berhasil disimpan", "data": seo})
}

func (h *SeoHandler) GetByReference(c *fiber.Ctx) error {
	refType := c.Query("type")
	refID := c.Query("id")

	if refType == "" || refID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Parameter type dan id wajib diisi"})
	}

	seo, err := h.seoUseCase.GetSeoByReference(c.Context(), refType, refID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "SEO Metadata tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"data": seo})
}
