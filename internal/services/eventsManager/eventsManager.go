package eventsmanager

import (
	"github.com/SaTeR151/EduMinin/internal/controller/rest/dto"
	"github.com/SaTeR151/EduMinin/internal/database/sqlite"
)

type EventsManagerService interface {
	GetEvents(limit string) ([]dto.Events, error)
	DeleteEvent(id string) error
	PostEvent(event dto.Events) error
}

type EventsManager struct {
	db sqlite.SQLiteManagerDB
}

func New(db sqlite.SQLiteManagerDB) *EventsManager {
	return &EventsManager{db: db}
}

func (em *EventsManager) GetEvents(limit string) ([]dto.Events, error) {
	var events []dto.Events
	events, err := em.db.GetEvents(limit)
	if err != nil {
		return nil, err
	}
	return events, nil
}
func (em *EventsManager) DeleteEvent(id string) error {
	return em.db.DeleteEvent(id)
}
func (em *EventsManager) PostEvent(event dto.Events) error {
	return em.db.SaveEvent(event)
}
