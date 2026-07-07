package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Table Model
type Table struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TableNumber int       `gorm:"uniqueIndex;not null" json:"table_number"`
	QRCodeURL   string    `gorm:"type:varchar(255)" json:"qr_code_url"`
	Status          string    `gorm:"type:varchar(50);default:'available'" json:"status"` // available, occupied
	NeedsAssistance bool      `gorm:"default:false" json:"needs_assistance"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Menu Model
type Menu struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	ImageURL    string    `gorm:"type:varchar(255)" json:"image_url"`
	Category    string    `gorm:"type:varchar(100)" json:"category"`
	IsAvailable bool      `gorm:"default:true" json:"is_available"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Order Model
type Order struct {
	ID              uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TableID         uuid.UUID   `gorm:"type:uuid;not null" json:"table_id"`
	Status          string      `gorm:"type:varchar(50);default:'pending'" json:"status"` // pending, preparing, served, paid, cancelled
	TotalAmount     float64     `gorm:"type:decimal(10,2);not null;default:0.00" json:"total_amount"`
	PaymentStatus   string      `gorm:"type:varchar(50);default:'unpaid'" json:"payment_status"` // unpaid, paid, failed
	PaymentIntentID string      `gorm:"type:varchar(255)" json:"payment_intent_id"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	Items           []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;" json:"items"`
}

// OrderItem Model
type OrderItem struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null" json:"order_id"`
	MenuID    uuid.UUID `gorm:"type:uuid;not null" json:"menu_id"`
	Menu      Menu      `gorm:"foreignKey:MenuID" json:"menu"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	UnitPrice float64   `gorm:"type:decimal(10,2);not null" json:"unit_price"`
	Subtotal  float64   `gorm:"type:decimal(10,2);not null" json:"subtotal"`
	Notes     string    `gorm:"type:text" json:"notes"`
	CreatedAt time.Time `json:"created_at"`
}

// BeforeCreate hooks for UUID generation in GORM if not relying purely on Postgres default
func (base *Table) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}
	return
}

func (base *Menu) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}
	return
}

func (base *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}
	return
}

func (base *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}
	return
}
