package repository

import (
	"myapp/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &orderRepository{db}
}

func (o *orderRepository) Create(order *domain.Order) error {
	return o.db.Create(order).Error
}

func (o *orderRepository) GetByID(id uuid.UUID) (*domain.Order, error) {
	var order domain.Order
	err := o.db.Preload("Items.Menu").First(&order, "id = ?", id).Error
	return &order, err
}

func (o *orderRepository) UpdateStatus(id uuid.UUID, status string) error {
	return o.db.Model(&domain.Order{}).Where("id = ?", id).Update("status", status).Error
}
