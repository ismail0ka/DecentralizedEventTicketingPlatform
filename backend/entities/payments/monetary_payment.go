package payments

import (
	"gorm.io/gorm"
)

type MonetaryPayment struct {
	gorm.Model
	Payment
	Amount          float32 `json: "amount"`
	Currency        string  `json: "currency"`
	TransactionId  string  `json: "transaction_id"`
	PaymentGateway string  `json: "payment_gateway"`
}
