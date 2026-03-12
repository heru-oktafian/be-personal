package handlers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	// 1. Parsing data teks dari multipart/form-data
	title := c.FormValue("title")
	description := c.FormValue("description")
	tags := c.FormValue("tags")
	isPublishedStr := c.FormValue("is_published")

	// Konversi string ke boolean
	isPublished := isPublishedStr == "true" || isPublishedStr == "1"

	// Validasi dasar
	if title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Kolom 'title' wajib diisi"})
	}

	// 2. Persiapan entitas
	project := entities.Project{
		Title:       title,
		Description: description,
		Tags:        tags,
		IsPublished: isPublished,
	}

	// 3. Menangani Upload File (Key: "image")
	file, err := c.FormFile("image")
	if err == nil {
		// Jika ada file yang diupload, kita proses

		// Buat folder jika belum ada (Mencegah error folder not found)
		uploadDir := "./public/uploads/projects"
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal membuat direktori upload"})
		}

		// Ambil ekstensi asli file (misal: .png, .jpg)
		ext := filepath.Ext(file.Filename)

		// Generate nama unik menggunakan UUID untuk mencegah bentrok nama file
		newFileName := uuid.New().String() + ext

		// Tentukan path penyimpanan fisik di server
		savePath := fmt.Sprintf("%s/%s", uploadDir, newFileName)

		// Simpan file ke server
		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan gambar"})
		}

		// Simpan alamat URL publiknya ke dalam database (Ini yang akan dibaca oleh Frontend)
		// Contoh hasil: "/uploads/projects/123e4567-e89b-12d3-a456-426614174000.png"
		project.ImageURL = fmt.Sprintf("/uploads/projects/%s", newFileName)
	}

	// 4. Eksekusi UseCase (Menambahkan ID, Slug, Timestamp, dan Simpan ke DB)
	if err := h.projectUseCase.CreateProject(c.Context(), &project); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan data proyek ke database"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Proyek portofolio berhasil dibuat beserta gambarnya",
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

	// 1. Ambil data proyek lama untuk memastikan eksistensi dan mengambil path gambar lama
	existingProject, err := h.projectUseCase.GetProjectByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Proyek tidak ditemukan"})
	}

	// 2. Parsing data teks dari multipart/form-data
	title := c.FormValue("title")
	description := c.FormValue("description")
	tags := c.FormValue("tags")
	isPublishedStr := c.FormValue("is_published")

	// Update field teks jika ada inputan baru
	if title != "" {
		existingProject.Title = title
	}
	if description != "" {
		existingProject.Description = description
	}
	if tags != "" {
		existingProject.Tags = tags
	}
	if isPublishedStr != "" {
		existingProject.IsPublished = isPublishedStr == "true" || isPublishedStr == "1"
	}

	// 3. Menangani Upload File Gambar Baru (Opsional)
	file, err := c.FormFile("image")
	if err == nil {
		// Jika form mengirimkan file gambar baru
		uploadDir := "./public/uploads/projects"
		os.MkdirAll(uploadDir, os.ModePerm)

		// Ekstraksi path file lama dan hapus dari server fisik untuk menghemat disk space
		if existingProject.ImageURL != "" {
			oldFilePath := fmt.Sprintf(".%s", existingProject.ImageURL)
			_ = os.Remove(oldFilePath) // Abaikan error jika file fisik sebelumnya sudah tidak ada
		}

		// Generate nama file baru
		ext := filepath.Ext(file.Filename)
		newFileName := uuid.New().String() + ext
		savePath := fmt.Sprintf("%s/%s", uploadDir, newFileName)

		// Simpan gambar baru
		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan gambar baru"})
		}

		// Timpa URL di database dengan gambar yang baru
		existingProject.ImageURL = fmt.Sprintf("/uploads/projects/%s", newFileName)
	}

	// Proteksi ganda agar ID tidak berubah dari payload berbahaya
	existingProject.ID = id

	// 4. Eksekusi UseCase untuk menyimpan pembaruan ke PostgreSQL
	if err := h.projectUseCase.UpdateProject(c.Context(), existingProject); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal memperbarui proyek"})
	}

	return c.JSON(fiber.Map{
		"message": "Proyek berhasil diperbarui",
		"data":    existingProject,
	})
}

func (h *ProjectHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	// 1. Ambil data proyek terlebih dahulu untuk mendapatkan path gambar
	existingProject, err := h.projectUseCase.GetProjectByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Proyek tidak ditemukan"})
	}

	// 2. Eksekusi UseCase untuk menghapus data dari database
	// (Pastikan fungsi delete di UseCase Anda bernama Delete atau DeleteProject sesuai yang Anda definisikan sebelumnya)
	if err := h.projectUseCase.DeleteProject(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus data proyek dari database"})
	}

	// 3. Hapus file gambar fisik dari server jika URL gambar tidak kosong
	if existingProject.ImageURL != "" {
		// Mengingat URL disimpan sebagai "/uploads/projects/namafile.ext"
		// dan folder fisiknya ada di "./public/uploads/projects/"
		filePath := fmt.Sprintf("./public%s", existingProject.ImageURL)

		// os.Remove akan menghapus file.
		// Kita abaikan error (_) untuk mencegah API gagal merespons sukses
		// jika kebetulan file fisiknya sudah terhapus secara manual sebelumnya.
		_ = os.Remove(filePath)
	}

	return c.JSON(fiber.Map{
		"message": "Proyek dan file gambarnya berhasil dihapus secara permanen",
	})
}
