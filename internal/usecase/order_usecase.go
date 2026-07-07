package usecase

import (
	"myapp/internal/domain"

	"github.com/google/uuid"
)

type orderUsecase struct {
	orderRepo domain.OrderRepository
	menuRepo  domain.MenuRepository
}

func NewOrderUsecase(or domain.OrderRepository, mr domain.MenuRepository) domain.OrderUsecase {
	return &orderUsecase{
		orderRepo: or,
		menuRepo:  mr,
	}
}

func (o *orderUsecase) CreateOrder(tableID uuid.UUID, items []domain.OrderItem) (*domain.Order, error) {
	var totalAmount float64

	// Calculate total amount based on current menu prices
	for i := range items {
		menu, err := o.menuRepo.GetByID(items[i].MenuID)
		if err != nil {
			return nil, err
		}
		items[i].UnitPrice = menu.Price
		items[i].Subtotal = menu.Price * float64(items[i].Quantity)
		totalAmount += items[i].Subtotal
	}

	order := &domain.Order{
		TableID:     tableID,
		Status:      "pending",
		TotalAmount: totalAmount,
		Items:       items,
	}

	err := o.orderRepo.Create(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *orderUsecase) GetOrder(id uuid.UUID) (*domain.Order, error) {
	return o.orderRepo.GetByID(id)
}

func (o *orderUsecase) UpdateOrderStatus(id uuid.UUID, status string) error {
	return o.orderRepo.UpdateStatus(id, status)
}
