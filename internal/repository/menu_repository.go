package repository

import (
	"myapp/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) domain.MenuRepository {
	return &menuRepository{db}
}

func (m *menuRepository) GetAll() ([]domain.Menu, error) {
	var menus []domain.Menu
	err := m.db.Where("is_available = ?", true).Find(&menus).Error
	return menus, err
}

func (m *menuRepository) GetByID(id uuid.UUID) (*domain.Menu, error) {
	var menu domain.Menu
	err := m.db.First(&menu, "id = ?", id).Error
	return &menu, err
}

func (m *menuRepository) Create(menu *domain.Menu) error {
	return m.db.Create(menu).Error
}
