package repositories

import (
	"gorm.io/gorm"
	"backend/entities/events"
)

type EventRepository struct {
	*BaseRepository[events.Event]
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{
		BaseRepository: NewBaseRepository[events.Event](db),
	};
}
