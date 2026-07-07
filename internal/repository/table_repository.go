package repository

import (
	"myapp/internal/domain"

	"gorm.io/gorm"
)

type tableRepository struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) domain.TableRepository {
	return &tableRepository{db}
}

func (t *tableRepository) GetByNumber(tableNumber int) (*domain.Table, error) {
	var table domain.Table
	err := t.db.First(&table, "table_number = ?", tableNumber).Error
	return &table, err
}

func (t *tableRepository) Create(table *domain.Table) error {
	return t.db.Create(table).Error
}

func (t *tableRepository) UpdateAssistance(tableNumber int, needsAssistance bool) error {
	return t.db.Model(&domain.Table{}).Where("table_number = ?", tableNumber).Update("needs_assistance", needsAssistance).Error
}
