package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/be-personal/internal/handlers"
	"github.com/heru-oktafian/be-personal/internal/repositories"
	"github.com/heru-oktafian/be-personal/internal/usecases"
	"github.com/heru-oktafian/be-personal/middleware"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(app *fiber.App, db *sqlx.DB) {
	// 1. Dependency Injection Setup
	// Repositories
	adminRepo := repositories.NewAdminRepository(db)
	projectRepo := repositories.NewProjectRepository(db)

	// UseCases
	authUseCase := usecases.NewAuthUseCase(adminRepo)
	projectUseCase := usecases.NewProjectUseCase(projectRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authUseCase)
	projectHandler := handlers.NewProjectHandler(projectUseCase)

	// --- DI: Repositories ---
	contactRepo := repositories.NewContactRepository(db)
	seoRepo := repositories.NewSeoRepository(db)
	settingRepo := repositories.NewSettingRepository(db)
	serviceRepo := repositories.NewServiceRepository(db)
	skillRepo := repositories.NewSkillRepository(db)
	// (adminRepo sudah ada, tinggal dipakai ulang)

	// --- DI: UseCases ---
	contactUseCase := usecases.NewContactUseCase(contactRepo)
	adminUseCase := usecases.NewAdminUseCase(adminRepo)
	seoUseCase := usecases.NewSeoUseCase(seoRepo)
	settingUC := usecases.NewSettingUseCase(settingRepo)
	serviceUC := usecases.NewServiceUseCase(serviceRepo)
	skillUC := usecases.NewSkillUseCase(skillRepo)

	// --- DI: Handlers ---
	contactHandler := handlers.NewContactHandler(contactUseCase)
	adminHandler := handlers.NewAdminHandler(adminUseCase)
	seoHandler := handlers.NewSeoHandler(seoUseCase)
	settingHandler := handlers.NewSettingHandler(settingUC)
	serviceHandler := handlers.NewServiceHandler(serviceUC)
	skillHandler := handlers.NewSkillHandler(skillUC)

	// --- DI: Home ---
	homeRepo := repositories.NewHomeRepository(db)
	homeUseCase := usecases.NewHomeUseCase(homeRepo)
	homeHandler := handlers.NewHomeHandler(homeUseCase)

	// 2. Routing Definition
	api := app.Group("/api/v1")

	// --- RUTE PUBLIK (Tanpa Auth) ---
	public := api.Group("/public")
	public.Get("/projects", projectHandler.GetAll)
	public.Get("/projects/:slug", projectHandler.GetBySlug)

	// --- RUTE PUBLIK ---
	public.Get("/seo", seoHandler.GetByReference)  // Rails get SEO dinamis (?type=Project&id=...)
	public.Post("/contact", contactHandler.Submit) // Visitor submit pesan
	public.Get("/home", homeHandler.GetHomepage)

	// Rute Auth
	api.Post("/auth/login", authHandler.Login)

	// --- RUTE ADMIN (Dilindungi JWT) ---
	admin := api.Group("/admin", middleware.Protected())
	admin.Post("/projects", projectHandler.Create)
	admin.Get("/projects", projectHandler.GetAll)
	admin.Get("/projects/:id", projectHandler.GetByID)
	admin.Put("/projects/:id", projectHandler.Update)
	admin.Delete("/projects/:id", projectHandler.Delete)
	admin.Get("/contacts", contactHandler.GetAll) // Baca daftar pesan
	admin.Get("/contacts/:id", contactHandler.GetByID)
	admin.Delete("/contacts/:id", contactHandler.Delete)
	admin.Get("/me", adminHandler.GetProfile) // Ambil data admin login
	admin.Post("/seo", seoHandler.Upsert)     // Simpan/update SEO

	// --- Settings ---
	admin.Get("/settings", settingHandler.GetAll)
	admin.Get("/settings/:key", settingHandler.GetByKey)  // Get Detail Setting
	admin.Post("/settings", settingHandler.Upsert)        // Add/Update Setting
	admin.Delete("/settings/:key", settingHandler.Delete) // Delete Setting

	// Services
	admin.Get("/services", serviceHandler.GetAll)
	admin.Post("/services", serviceHandler.Create)
	admin.Get("/services/:id", serviceHandler.GetByID)
	admin.Put("/services/:id", serviceHandler.Update)
	admin.Delete("/services/:id", serviceHandler.Delete)

	// --- Skills ---
	admin.Get("/skills", skillHandler.GetAll)
	admin.Get("/skills/:id", skillHandler.GetByID)   // Get Detail Skill
	admin.Post("/skills", skillHandler.Create)       // Add New Skill
	admin.Put("/skills/:id", skillHandler.Update)    // Update Skill
	admin.Delete("/skills/:id", skillHandler.Delete) // Delete Skill

}
