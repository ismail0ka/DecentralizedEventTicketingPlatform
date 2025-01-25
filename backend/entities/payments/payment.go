package payments

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserID        int64         `json:"user_id"`
	EventID       int64         `json:"event_id"`
	TicketsCount  float32       `json:"tickets_count"`
	PaymentDate   string        `json:"payment_date"`
	PaymentMethod PaymentMethod `json:"payment_method"`
}
