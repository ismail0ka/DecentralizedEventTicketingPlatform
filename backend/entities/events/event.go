package events

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	Date             string  `json:"date"`
	Location         string  `json:"location"`
	OrganizerID      int64   `json:"organizer_id"`
	AvailableTickets int32   `json:"available_tickets"`
	Price            float32 `json:"price"`
	ImageURL         string  `json:"image_url"`
}
