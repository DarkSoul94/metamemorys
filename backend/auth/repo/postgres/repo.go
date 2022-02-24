package postgres

import (
	"github.com/DarkSoul94/metamemorys_backend/auth"
	"github.com/DarkSoul94/metamemorys_backend/models"
	"gorm.io/gorm"
)

type postgresRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) auth.AuthRepo {
	return &postgresRepo{
		db: db,
	}
}

func (r *postgresRepo) CreateUser(user *models.User) error {
	return r.db.Create(&user).Error
}

func (r *postgresRepo) GetUser(email string) (*models.User, error) {
	var res *models.User
	tx := r.db.Model(&models.User{}).Where(&models.User{Email: email}).First(&res)
	return res, tx.Error
}
