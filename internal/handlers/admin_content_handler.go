package handlers

import (
	"time"

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

// Get Detail Setting
func (h *SettingHandler) GetByKey(c *fiber.Ctx) error {
	key := c.Params("key")
	res, err := h.uc.GetSettingByKey(c.Context(), key)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pengaturan tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"data": res})
}

// Add New Setting & Update Setting
// (Karena berbasis Key-Value, Create dan Update menggunakan metode Upsert yang sama)
func (h *SettingHandler) Upsert(c *fiber.Ctx) error {
	var payload entities.SiteSetting
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format request tidak valid"})
	}

	if payload.Key == "" || payload.Value == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Key dan Value wajib diisi"})
	}

	// WAJIB DITAMBAHKAN: Set waktu pembaruan agar tidak ditolak database
	payload.UpdatedAt = time.Now()

	if err := h.uc.SaveSetting(c.Context(), &payload); err != nil {
		// BUKA BLOKIR ERROR: Kita cetak err.Error() agar tahu persis apa yang salah di SQL
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Gagal menyimpan pengaturan",
			"detail": err.Error(), // <--- Ini kunci utamanya
		})
	}

	return c.JSON(fiber.Map{
		"message": "Pengaturan berhasil disimpan",
		"data":    payload,
	})
}

// Delete Setting
func (h *SettingHandler) Delete(c *fiber.Ctx) error {
	key := c.Params("key")
	if err := h.uc.DeleteSetting(c.Context(), key); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus pengaturan"})
	}
	return c.JSON(fiber.Map{"message": "Pengaturan berhasil dihapus"})
}
