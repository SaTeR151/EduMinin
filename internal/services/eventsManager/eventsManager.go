package eventsmanager

import (
	"github.com/SaTeR151/EduMinin/internal/database/sqlite"
)

type EventsManagerService interface {
}

type EventsManager struct {
	db sqlite.SQLiteManagerDB
}

func New(db sqlite.SQLiteManagerDB) *EventsManager {
	return &EventsManager{db: db}
}
