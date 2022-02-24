package http

import "github.com/DarkSoul94/metamemorys_backend/models"

func (h *AuthHandler) toModelUser(user UserAuth) *models.User {
	return &models.User{
		Email:    user.Email,
		PassHash: user.Pass,
	}
}
