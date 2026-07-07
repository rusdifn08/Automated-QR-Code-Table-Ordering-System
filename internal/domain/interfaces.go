package domain

import "github.com/google/uuid"

// MenuRepository defines the interface for Menu data access
type MenuRepository interface {
	GetAll() ([]Menu, error)
	GetByID(id uuid.UUID) (*Menu, error)
	Create(menu *Menu) error
}

// MenuUsecase defines the interface for Menu business logic
type MenuUsecase interface {
	FetchAll() ([]Menu, error)
}

// OrderRepository defines the interface for Order data access
type OrderRepository interface {
	GetAll() ([]Order, error)
	Create(order *Order) error
	GetByID(id uuid.UUID) (*Order, error)
	UpdateStatus(id uuid.UUID, status string) error
}

// OrderUsecase defines the interface for Order business logic
type OrderUsecase interface {
	GetAllOrders() ([]Order, error)
	CreateOrder(tableID uuid.UUID, items []OrderItem) (*Order, error)
	GetOrder(id uuid.UUID) (*Order, error)
	UpdateOrderStatus(id uuid.UUID, status string) error
}

// TableRepository defines the interface for Table data access
type TableRepository interface {
	GetByNumber(tableNumber int) (*Table, error)
	Create(table *Table) error
	UpdateAssistance(tableNumber int, needsAssistance bool) error
}

// TableUsecase defines the interface for Table business logic
type TableUsecase interface {
	GetTableByNumber(tableNumber int) (*Table, error)
	CallWaiter(tableNumber int) error
	ResolveAssistance(tableNumber int) error
}
