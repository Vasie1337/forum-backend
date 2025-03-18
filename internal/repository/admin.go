package repository

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return AdminRepository{db}
}

func (r *AdminRepository) GetByUsername(username string) (*models.Admin, error) {
	var admin models.Admin
	err := r.db.Where("username = ?", username).First(&admin).Error
	return &admin, err
}

func (r *AdminRepository) Create(admin *models.Admin) error {
	return r.db.Create(admin).Error
}

func (r *AdminRepository) Update(admin *models.Admin) error {
	return r.db.Save(admin).Error
}

func (r *AdminRepository) Delete(admin *models.Admin) error {
	return r.db.Delete(admin).Error
}
