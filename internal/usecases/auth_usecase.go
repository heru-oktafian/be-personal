package usecases

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/heru-oktafian/be-personal/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	Login(ctx context.Context, email, password string) (string, error)
}

type authUseCase struct {
	adminRepo repositories.AdminRepository
}

func NewAuthUseCase(adminRepo repositories.AdminRepository) AuthUseCase {
	return &authUseCase{adminRepo: adminRepo}
}

func (u *authUseCase) Login(ctx context.Context, email, password string) (string, error) {
	// 1. Cek keberadaan admin berdasarkan email
	admin, err := u.adminRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("kredensial tidak valid")
	}

	// 2. Verifikasi hash password
	err = bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("kredensial tidak valid")
	}

	// 3. Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": admin.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Berlaku 24 jam
	})

	secret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", errors.New("gagal membuat token autentikasi")
	}

	return tokenString, nil
}
