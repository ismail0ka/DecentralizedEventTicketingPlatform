package payments

type Payment struct {
	id             int64         `json: "id"`
	user_id        int64         `json: "user_id"`
	event_id       int64         `json: "event_id"`
	tickets_count  float32       `json: "tickets_count"`
	payment_date   string        `json: "payment_date"`
	payment_method PaymentMethod `json: "payment_method"`
}
