package newsmanager

import (
	"github.com/SaTeR151/EduMinin/internal/controller/rest/dto"
	"github.com/SaTeR151/EduMinin/internal/database/sqlite"
)

type NewsManagerService interface {
	GetNews(limit string) ([]dto.News, error)
	DeleteNews(id string) error
	PostNews(news dto.News) error
}

type NewsManager struct {
	db sqlite.SQLiteManagerDB
}

func New(db sqlite.SQLiteManagerDB) *NewsManager {
	return &NewsManager{db: db}
}

func (nm *NewsManager) GetNews(limit string) ([]dto.News, error) {
	var news []dto.News
	news, err := nm.db.GetNews(limit)
	if err != nil {
		return nil, err
	}
	return news, nil
}
func (nm *NewsManager) DeleteNews(id string) error {
	return nm.db.DeleteNews(id)
}
func (nm *NewsManager) PostNews(news dto.News) error {
	return nm.db.SaveNews(news)
}
