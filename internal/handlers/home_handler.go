package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/be-personal/internal/usecases"
)

type HomeHandler struct {
	homeUseCase usecases.HomeUseCase
}

func NewHomeHandler(homeUseCase usecases.HomeUseCase) *HomeHandler {
	return &HomeHandler{homeUseCase: homeUseCase}
}

func (h *HomeHandler) GetHomepage(c *fiber.Ctx) error {
	data, err := h.homeUseCase.GetHomeData(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal merender data homepage",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Sukses",
		"data":    data,
	})
}
