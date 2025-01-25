package events

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	EventID      int64        `json:"event_id"`
	UserID       int64        `json:"user_id"`
	EventStatus  EventStatus  `json:"event_status"`
	QRCode       string       `json:"qr_code"`
	TicketStatus TicketStatus `json:"ticket_status"`
}
