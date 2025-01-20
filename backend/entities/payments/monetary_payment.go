package payments

type MonetaryPayment struct {
	Payment
	amount          float32 `json: "amount"`
	currency        string  `json: "currency"`
	transaction_id  string  `json: "transaction_id"`
	payment_gateway string  `json: "payment_gateway"`
}
