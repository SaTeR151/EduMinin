package reviewsmanager

import (
	"github.com/SaTeR151/EduMinin/internal/controller/rest/dto"
	"github.com/SaTeR151/EduMinin/internal/database/sqlite"
)

type ReviewsManagerService interface {
}

type ReviewsManager struct {
	db sqlite.SQLiteManagerDB
}

func New(db sqlite.SQLiteManagerDB) *ReviewsManager {
	return &ReviewsManager{db: db}
}

func (rm *ReviewsManager) GetReviews(limit string) ([]dto.Reviews, error) {
	var reviews []dto.Reviews
	reviews, err := rm.db.GetReviews(limit)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}
func (rm *ReviewsManager) DeleteReview(id string) error {
	return rm.db.DeleteReview(id)
}
func (rm *ReviewsManager) PostReview(review dto.Reviews) error {
	return rm.db.SaveReview(review)
}
