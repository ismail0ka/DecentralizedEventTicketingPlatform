package payments

import (
	"gorm.io/gorm"
)

type CryptoPayment struct {
	gorm.Model
	Payment
	Amount           float32 `json: "amount"`
	Currency         string  `json: "currency"`
	Transaction_hash string  `json: "transaction_hash"`
	WalletAddress   string  `json: "wallet_address"`
}
