package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/heru-oktafian/be-personal/internal/usecases"
)

// --- SETTING HANDLER ---
type SettingHandler struct{ uc usecases.SettingUseCase }

func NewSettingHandler(uc usecases.SettingUseCase) *SettingHandler { return &SettingHandler{uc} }

func (h *SettingHandler) GetAll(c *fiber.Ctx) error {
	res, err := h.uc.GetSettings(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": res})
}

func (h *SettingHandler) Upsert(c *fiber.Ctx) error {
	var payload entities.SiteSetting
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid format"})
	}
	if err := h.uc.SaveSetting(c.Context(), &payload); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Setting saved", "data": payload})
}
