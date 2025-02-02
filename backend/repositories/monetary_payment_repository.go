package repositories

import (
	"gorm.io/gorm"
	"backend/entities/payments"
)

type MonetaryPaymentRepository struct {
	*BaseRepository[payments.MonetaryPayment]
}

func NewMonetaryPaymentRepository(db *gorm.DB) *MonetaryPaymentRepository {
	return &MonetaryPaymentRepository{
		BaseRepository: NewBaseRepository[payments.MonetaryPayment](db),
	};
}