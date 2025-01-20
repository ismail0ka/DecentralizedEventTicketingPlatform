package payments

type CryptoPayment struct {
	Payment
	amount           float32 `json: "amount"`
	currency         string  `json: "currency"`
	transaction_hash string  `json: "transaction_hash"`
	wallet_address   string  `json: "wallet_address"`
}
