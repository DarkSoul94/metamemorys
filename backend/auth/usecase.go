package auth

import (
	"context"

	"github.com/DarkSoul94/metamemorys_backend/models"
)

type AuthUC interface {
	Registration(email, pass, name string) (string, error)
	Authorization(user *models.User) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
}
