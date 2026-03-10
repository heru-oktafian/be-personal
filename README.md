# Heru Oktafian - Personal Website CMS API

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![Fiber](https://img.shields.io/badge/Fiber-v2-black?style=for-the-badge&logo=gofiber)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-316192?style=for-the-badge&logo=postgresql)
![JWT](https://img.shields.io/badge/JWT-Auth-black?style=for-the-badge&logo=jsonwebtokens)

API Backend tangguh yang dirancang khusus untuk mengelola konten (CMS) pada heruoktafian.com. Dibangun menggunakan **Golang**, **Fiber Web Framework**, dan **PostgreSQL**, API ini berfungsi sebagai *headless CMS* yang melayani halaman publik (Ruby on Rails) dan dasbor admin (React SPA).

## 🏗️ Arsitektur (Clean Architecture)

Proyek ini secara ketat mengimplementasikan **Clean Architecture** berlapis untuk memastikan pemisahan tanggung jawab (*Separation of Concerns*), kemudahan *testing*, dan skalabilitas jangka panjang:

1. **Entities**: Struktur data domain murni (tidak bergantung pada *framework*).
2. **Repositories**: Layer interaksi *database* (PostgreSQL via `sqlx`).
3. **UseCases**: Pusat logika bisnis dan aturan aplikasi.
4. **Handlers (Delivery)**: Layer HTTP *controllers* menggunakan Fiber untuk mengatur *request* dan *response* JSON.

## 🚀 Fitur Utama

- **Decoupled Architecture**: API murni (JSON) tanpa campur tangan *rendering* UI.
- **Secure Admin Authentication**: Perlindungan *endpoint* admin menggunakan JWT (JSON Web Token) dan *password hashing* (`bcrypt`).
- **Portfolio/Project Management**: CRUD lengkap untuk mengelola proyek portofolio.
- **Dynamic SEO Metadata**: Dukungan *endpoint* untuk pengaturan metadata SEO per halaman/entitas.
- **High Performance**: Memanfaatkan Fiber framework yang terkenal akan kecepatannya di ekosistem Go.

## 📂 Struktur Direktori

```text
cms-backend/
├── cmd/
│   └── api/
│       └── main.go           # Titik masuk aplikasi
├── config/
│   └── database.go           # Konfigurasi koneksi PostgreSQL
├── internal/
│   ├── entities/             # Model domain murni
│   ├── handlers/             # HTTP Controllers (Fiber)
│   ├── repositories/         # Akses data PostgreSQL (sqlx)
│   └── usecases/             # Logika bisnis
├── middleware/
│   └── jwt.go                # Middleware proteksi rute admin
├── routes/
│   └── api.go                # Registrasi endpoint
├── .env.example
├── go.mod
└── go.sum