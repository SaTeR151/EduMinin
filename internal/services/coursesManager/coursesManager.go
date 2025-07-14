package coursesmanager

import (
	"github.com/SaTeR151/EduMinin/internal/controller/rest/dto"
	"github.com/SaTeR151/EduMinin/internal/database/sqlite"
)

type CoursesManagerService interface {
	GetCourses(limit string) ([]dto.Course, error)
	DeleteCourse(id string) error
	PostCourse(course dto.Course) error
}

type CoursesManager struct {
	db sqlite.SQLiteManagerDB
}

func New(db sqlite.SQLiteManagerDB) *CoursesManager {
	return &CoursesManager{db: db}
}

func (cm *CoursesManager) GetCourses(limit string) ([]dto.Course, error) {
	var courses []dto.Course
	courses, err := cm.db.GetCourses(limit)
	if err != nil {
		return nil, err
	}
	return courses, nil
}
func (cm *CoursesManager) DeleteCourse(id string) error {
	return cm.db.DeleteCourse(id)
}
func (cm *CoursesManager) PostCourse(course dto.Course) error {
	return cm.db.SaveCourse(course)
}
