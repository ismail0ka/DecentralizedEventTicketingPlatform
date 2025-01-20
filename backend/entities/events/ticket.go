package events

type Ticket struct {
	id            int64        `json: "id"`
	event_id      int64        `json: "event_id"`
	user_id       int64        `json: "user_id"`
	event_status  EventStatus  `json: "event_status"`
	qr_code       string       `json: "qr_code"`
	ticket_status TicketStatus `json: "ticket_status"`
}
