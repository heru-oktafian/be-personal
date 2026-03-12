package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

// Get Detail Skill
func (h *SkillHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.uc.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Keahlian tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"data": res})
}

// Add New Skill
func (h *SkillHandler) Create(c *fiber.Ctx) error {
	name := c.FormValue("name")
	category := c.FormValue("category")

	// Konversi string ke integer
	percentage, _ := strconv.Atoi(c.FormValue("percentage"))
	orderNum, _ := strconv.Atoi(c.FormValue("order_num"))

	if name == "" || category == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Nama dan Kategori wajib diisi"})
	}

	skill := entities.Skill{
		Name:       name,
		Category:   category,
		Percentage: percentage,
		OrderNum:   orderNum,
	}

	// Menangani Upload Icon (Opsional, karena tidak semua skill mungkin punya icon)
	file, err := c.FormFile("icon")
	if err == nil {
		uploadDir := "./public/uploads/skills"
		os.MkdirAll(uploadDir, os.ModePerm)

		ext := filepath.Ext(file.Filename)
		newFileName := uuid.New().String() + ext
		savePath := fmt.Sprintf("%s/%s", uploadDir, newFileName)

		if err := c.SaveFile(file, savePath); err == nil {
			skill.IconURL = fmt.Sprintf("/uploads/skills/%s", newFileName)
		}
	}

	if err := h.uc.Create(c.Context(), &skill); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan keahlian"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Keahlian berhasil ditambahkan",
		"data":    skill,
	})
}

// Update Skill
func (h *SkillHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	existingSkill, err := h.uc.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Keahlian tidak ditemukan"})
	}

	// Parsing input teks
	if name := c.FormValue("name"); name != "" {
		existingSkill.Name = name
	}
	if category := c.FormValue("category"); category != "" {
		existingSkill.Category = category
	}
	if percentageStr := c.FormValue("percentage"); percentageStr != "" {
		existingSkill.Percentage, _ = strconv.Atoi(percentageStr)
	}
	if orderNumStr := c.FormValue("order_num"); orderNumStr != "" {
		existingSkill.OrderNum, _ = strconv.Atoi(orderNumStr)
	}

	// Menangani Update Icon
	file, err := c.FormFile("icon")
	if err == nil {
		uploadDir := "./public/uploads/skills"
		os.MkdirAll(uploadDir, os.ModePerm)

		// Hapus icon lama
		if existingSkill.IconURL != "" {
			oldFilePath := fmt.Sprintf("./public%s", existingSkill.IconURL)
			_ = os.Remove(oldFilePath)
		}

		ext := filepath.Ext(file.Filename)
		newFileName := uuid.New().String() + ext
		savePath := fmt.Sprintf("%s/%s", uploadDir, newFileName)

		if err := c.SaveFile(file, savePath); err == nil {
			existingSkill.IconURL = fmt.Sprintf("/uploads/skills/%s", newFileName)
		}
	}

	if err := h.uc.Update(c.Context(), existingSkill); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal memperbarui keahlian"})
	}

	return c.JSON(fiber.Map{
		"message": "Keahlian berhasil diperbarui",
		"data":    existingSkill,
	})
}

// Delete Skill
func (h *SkillHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	existingSkill, err := h.uc.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Keahlian tidak ditemukan"})
	}

	if err := h.uc.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus data keahlian dari database"})
	}

	// Hapus file icon fisik dari server
	if existingSkill.IconURL != "" {
		filePath := fmt.Sprintf("./public%s", existingSkill.IconURL)
		_ = os.Remove(filePath)
	}

	return c.JSON(fiber.Map{"message": "Keahlian dan icon berhasil dihapus"})
}
