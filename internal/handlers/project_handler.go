package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/heru-oktafian/be-personal/internal/usecases"
)

type ProjectHandler struct {
	projectUseCase usecases.ProjectUseCase
}

// NewProjectHandler menginisialisasi dan mengembalikan instance baru dari ProjectHandler.
func NewProjectHandler(projectUseCase usecases.ProjectUseCase) *ProjectHandler {
	return &ProjectHandler{projectUseCase: projectUseCase}
}

// GetAll menangani HTTP request untuk mengambil semua data proyek yang tersedia.
// Mengembalikan daftar proyek dalam format JSON atau pesan error jika gagal.
func (h *ProjectHandler) GetAll(c *fiber.Ctx) error {
	projects, err := h.projectUseCase.GetAllProjects(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data proyek",
		})
	}

	return c.JSON(fiber.Map{
		"data": projects,
	})
}

// GetBySlug menangani HTTP request untuk mengambil data spesifik suatu proyek berdasarkan parameter slug.
// Mengembalikan detail proyek dalam format JSON jika ditemukan, atau pesan error (404 Not Found) jika tidak ada.
func (h *ProjectHandler) GetBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	project, err := h.projectUseCase.GetProjectBySlug(c.Context(), slug)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Proyek tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"data": project,
	})
}

// Create menangani HTTP request untuk menambahkan atau membuat data proyek baru.
// Fungsi ini mem-parsing body request dan menyimpannya melalui usecase.
// Mengembalikan status 201 Created beserta data proyek jika sukses, atau pesan error jika format tidak valid atau gagal menyimpan.
func (h *ProjectHandler) Create(c *fiber.Ctx) error {
	var project entities.Project
	if err := c.BodyParser(&project); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Format request tidak valid",
		})
	}

	if err := h.projectUseCase.CreateProject(c.Context(), &project); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal menyimpan proyek",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Proyek berhasil dibuat",
		"data":    project,
	})
}

func (h *ProjectHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	project, err := h.projectUseCase.GetProjectByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Proyek tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"data": project})
}

func (h *ProjectHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	// Ambil data lama untuk memastikan eksistensi dan mempertahankan created_at
	existingProject, err := h.projectUseCase.GetProjectByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Proyek tidak ditemukan"})
	}

	if err := c.BodyParser(existingProject); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format request tidak valid"})
	}

	existingProject.ID = id // Pastikan ID tidak tertimpa payload

	if err := h.projectUseCase.UpdateProject(c.Context(), existingProject); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal memperbarui proyek"})
	}

	return c.JSON(fiber.Map{"message": "Proyek berhasil diperbarui", "data": existingProject})
}

func (h *ProjectHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.projectUseCase.DeleteProject(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus proyek"})
	}
	return c.JSON(fiber.Map{"message": "Proyek berhasil dihapus"})
}
