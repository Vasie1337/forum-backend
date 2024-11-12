package repository

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type KeyRepository struct {
	db *gorm.DB
}

func NewKeyRepository(db *gorm.DB) *KeyRepository {
	return &KeyRepository{db: db}
}

func (r *KeyRepository) GetByID(id int) (*models.Key, error) {
	var key models.Key
	err := r.db.Where("id = ?", id).First(&key).Error
	return &key, err
}

func (r *KeyRepository) GetByValue(value string) (*models.Key, error) {
	var key models.Key
	err := r.db.Where("value = ?", value).First(&key).Error
	return &key, err
}

func (r *KeyRepository) Create(key *models.Key) error {
	return r.db.Create(key).Error
}

func (r *KeyRepository) Update(key *models.Key) error {
	return r.db.Save(key).Error
}
