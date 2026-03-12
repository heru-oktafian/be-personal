package usecases

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/heru-oktafian/be-personal/internal/entities"
	"github.com/heru-oktafian/be-personal/internal/repositories"
)

type ProjectUseCase interface {
	GetAllProjects(ctx context.Context) ([]entities.Project, error)
	GetProjectBySlug(ctx context.Context, slug string) (*entities.Project, error)
	CreateProject(ctx context.Context, project *entities.Project) error
	GetProjectByID(ctx context.Context, id string) (*entities.Project, error)
	UpdateProject(ctx context.Context, project *entities.Project) error
	DeleteProject(ctx context.Context, id string) error
}

type projectUseCase struct {
	projectRepo repositories.ProjectRepository
}

func NewProjectUseCase(projectRepo repositories.ProjectRepository) ProjectUseCase {
	return &projectUseCase{projectRepo: projectRepo}
}

func (u *projectUseCase) GetAllProjects(ctx context.Context) ([]entities.Project, error) {
	// Logika tambahan bisa disisipkan di sini jika diperlukan (misal: filter data tertentu)
	return u.projectRepo.FetchAll(ctx)
}

func (u *projectUseCase) GetProjectBySlug(ctx context.Context, slug string) (*entities.Project, error) {
	return u.projectRepo.GetBySlug(ctx, slug)
}

func (u *projectUseCase) CreateProject(ctx context.Context, project *entities.Project) error {
	// 1. Generate UUID
	project.ID = uuid.New().String()

	// 2. Auto-generate Slug jika kosong (Logika SEO)
	if project.Slug == "" {
		project.Slug = generateSlug(project.Title)
	}

	// 3. Set Timestamps
	now := time.Now()
	project.CreatedAt = now
	project.UpdatedAt = now

	// 4. Lempar ke Repository
	return u.projectRepo.Create(ctx, project)
}

// Fungsi helper (private) untuk membersihkan judul menjadi URL slug
func generateSlug(title string) string {
	// Ubah ke lowercase
	slug := strings.ToLower(title)
	// Hanya sisakan huruf, angka, dan spasi
	re := regexp.MustCompile(`[^a-z0-9\s]+`)
	slug = re.ReplaceAllString(slug, "")
	// Ganti spasi menjadi strip (-)
	reSpace := regexp.MustCompile(`\s+`)
	slug = reSpace.ReplaceAllString(slug, "-")

	return slug
}

func (u *projectUseCase) GetProjectByID(ctx context.Context, id string) (*entities.Project, error) {
	return u.projectRepo.GetByID(ctx, id)
}

func (u *projectUseCase) UpdateProject(ctx context.Context, project *entities.Project) error {
	project.UpdatedAt = time.Now() // Selalu perbarui timestamp
	return u.projectRepo.Update(ctx, project)
}

func (u *projectUseCase) DeleteProject(ctx context.Context, id string) error {
	return u.projectRepo.Delete(ctx, id)
}
