package services

import (
	"github.com/SaTeR151/EduMinin/internal/database/sqlite"
	coursesmanager "github.com/SaTeR151/EduMinin/internal/services/coursesManager"
	eventsmanager "github.com/SaTeR151/EduMinin/internal/services/eventsManager"
	newsmanager "github.com/SaTeR151/EduMinin/internal/services/newsManager"
	reviewsmanager "github.com/SaTeR151/EduMinin/internal/services/reviewsManager"
)

type Services struct {
	EventsManager  eventsmanager.EventsManagerService
	NewsManager    newsmanager.NewsManagerService
	ReviewsManager reviewsmanager.ReviewsManagerService
	CoursesManager coursesmanager.CoursesManagerService
}

func New(db sqlite.SQLiteManagerDB) *Services {
	var services Services
	services.EventsManager = eventsmanager.New(db)
	services.NewsManager = newsmanager.New(db)
	services.ReviewsManager = reviewsmanager.New(db)
	services.CoursesManager = coursesmanager.New(db)
	return &services
}
