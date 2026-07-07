package repository

import (
	"myapp/internal/domain"
	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) domain.AdminRepository {
	return &adminRepository{db}
}

func (a *adminRepository) GetByUsername(username string) (*domain.Admin, error) {
	var admin domain.Admin
	err := a.db.Where("username = ?", username).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (a *adminRepository) Create(admin *domain.Admin) error {
	return a.db.Create(admin).Error
}
