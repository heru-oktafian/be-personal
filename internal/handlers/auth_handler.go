package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/be-personal/internal/usecases"
)

type AuthHandler struct {
	authUseCase usecases.AuthUseCase
}

func NewAuthHandler(authUseCase usecases.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Format request tidak valid",
		})
	}

	token, err := h.authUseCase.Login(c.Context(), input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login berhasil",
		"token":   token,
	})
}
