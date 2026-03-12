package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/be-personal/internal/usecases"
)

type AdminHandler struct {
	adminUseCase usecases.AdminUseCase
}

func NewAdminHandler(adminUseCase usecases.AdminUseCase) *AdminHandler {
	return &AdminHandler{adminUseCase: adminUseCase}
}

func (h *AdminHandler) GetProfile(c *fiber.Ctx) error {
	// user_id diambil dari c.Locals yang diset oleh middleware/jwt.go
	userID := c.Locals("user_id").(string)

	admin, err := h.adminUseCase.GetProfile(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Admin tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"data": admin})
}
