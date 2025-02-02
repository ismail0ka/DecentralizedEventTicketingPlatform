package repositories

import (
	"gorm.io/gorm"
	"backend/entities/payments"
)

type CryptoPaymentRepository struct {
	*BaseRepository[payments.CryptoPayment]
}

//TO DO : Add Blockchain Client API to interact with the blockchain

func NewCryptoPaymentRepository(db *gorm.DB) *CryptoPaymentRepository {
	return &CryptoPaymentRepository{
		BaseRepository: &BaseRepository[payments.CryptoPayment]{db: db},
	};
}