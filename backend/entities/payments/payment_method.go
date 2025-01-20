package payments

type PaymentMethod int

const (
	monetary_payment PaymentMethod = iota
	crypto_payment
)

func (pm PaymentMethod) String() string{
	return [...]string{"Monetary", "Crypto"}[pm]
}