package auth

import "github.com/DarkSoul94/metamemorys_backend/models"

type AuthRepo interface {
	CreateUser(user *models.User) error
	GetUser(email string) (*models.User, error)
}
