package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/SaTeR151/EduMinin/internal/apperror"
	"github.com/SaTeR151/EduMinin/internal/config"
	"github.com/SaTeR151/EduMinin/internal/controller/rest/dto"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "modernc.org/sqlite"
)

type SQLiteManagerDB interface {
	MigrationUP() error
	MigrationDown() error

	SaveCourse(course dto.Course) error
	DeleteCourse(id string) error
	GetCourses(limit string) ([]dto.Course, error)

	SaveNews(news dto.News) error
	DeleteNews(id string) error
	GetNews(limit string) ([]dto.News, error)

	SaveReview(review dto.Reviews) error
	DeleteReview(id string) error
	GetReviews(limit string) ([]dto.Reviews, error)
}

type SQLiteManager struct {
	db *sql.DB
}

func Open(sqliteConfig config.SQLiteConfig) (*SQLiteManager, func() error, error) {
	var sqliteManger SQLiteManager
	var err error
	sqliteManger.db, err = sql.Open("sqlite", "./data.db")
	if err != nil {
		return &sqliteManger, nil, err
	}
	_, err = sqliteManger.db.Exec("")
	if err != nil {
		return &sqliteManger, nil, err
	}
	return &sqliteManger, sqliteManger.db.Close, nil
}

func (db *SQLiteManager) MigrationUP() error {
	driver, err := sqlite3.WithInstance(db.db, &sqlite3.Config{})
	if err != nil {
		return err
	}

	// Создаем экземпляр мигратора
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "sqlite3", driver)
	if err != nil {
		return err
	}

	// Применяем миграции
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {

		return err
	}
	return nil
}

func (db *SQLiteManager) MigrationDown() error {
	driver, err := sqlite3.WithInstance(db.db, &sqlite3.Config{})
	if err != nil {
		return err
	}

	// Создаем экземпляр мигратора
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "sqlite3", driver)
	if err != nil {
		log.Fatal(err)
	}

	// Применяем миграции
	err = m.Down()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func (db *SQLiteManager) SaveCourse(course dto.Course) error {
	_, err := db.db.Exec("INSERT INTO courses (title, description, academic_hours, price, image) values (:title, :description, :academic_hours, :price, :image)",
		sql.Named("title", course.Title),
		sql.Named("description", course.Description),
		sql.Named("academic_hours", course.AcademicHours),
		sql.Named("price", course.Price),
		sql.Named("image", course.Image),
	)
	return err
}

func (db *SQLiteManager) DeleteCourse(id string) error {
	res, err := db.db.Exec("DELETE FROM courses WHERE id = :id", sql.Named("id", id))
	rows, err := res.RowsAffected()
	if err != nil {
		return nil
	}
	if rows == 0 {
		return apperror.ErrDataNotFound
	}
	return err
}

func (db *SQLiteManager) GetCourses(limit string) ([]dto.Course, error) {
	query := "SELECT id, title, description, academic_hours, price, image FROM courses"
	if limit != "" {
		query += fmt.Sprintf(" LIMIT %s", limit)
	}
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var news []dto.Course
	for rows.Next() {
		var n dto.Course
		err := rows.Scan(&n.Id, &n.Title, &n.Description, &n.AcademicHours, &n.Price, &n.Image)
		if err != nil {
			return nil, err
		}
		news = append(news, n)
	}
	return news, nil
}

func (db *SQLiteManager) SaveNews(news dto.News) error {
	_, err := db.db.Exec("INSERT INTO news (title, content, date, image_path) values (:title, :content, :date, :image_path)",
		sql.Named("title", news.Title),
		sql.Named("content", news.Text),
		sql.Named("date", news.Date),
		sql.Named("image_path", news.Image),
	)
	return err
}

func (db *SQLiteManager) DeleteNews(id string) error {
	res, err := db.db.Exec("DELETE FROM news WHERE id = :id", sql.Named("id", id))
	rows, err := res.RowsAffected()
	if err != nil {
		return nil
	}
	if rows == 0 {
		return apperror.ErrDataNotFound
	}
	return err
}

func (db *SQLiteManager) GetNews(limit string) ([]dto.News, error) {
	query := "SELECT id, title, content, date, image_path image FROM news"
	if limit != "" {
		query += fmt.Sprintf(" LIMIT %s", limit)
	}
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var news []dto.News
	for rows.Next() {
		var n dto.News
		err := rows.Scan(&n.Id, &n.Title, &n.Text, &n.Date, &n.Image)
		if err != nil {
			return nil, err
		}
		news = append(news, n)
	}
	return news, nil
}

func (db *SQLiteManager) SaveReview(reviews dto.Reviews) error {
	_, err := db.db.Exec("INSERT INTO reviews (author_name, date, review_text, photo_path) values (:author_name, :date, :review_text, :photo_path)",
		sql.Named("author_name", reviews.Author),
		sql.Named("date", reviews.Date),
		sql.Named("review_text", reviews.Text),
		sql.Named("photo_path", reviews.Photo),
	)
	return err
}

func (db *SQLiteManager) DeleteReview(id string) error {
	res, err := db.db.Exec("DELETE FROM reviews WHERE id = :id", sql.Named("id", id))
	rows, err := res.RowsAffected()
	if err != nil {
		return nil
	}
	if rows == 0 {
		return apperror.ErrDataNotFound
	}
	return err
}

func (db *SQLiteManager) GetReviews(limit string) ([]dto.Reviews, error) {
	query := "SELECT id, author_name, date, review_text, photo_path image FROM reviews"
	if limit != "" {
		query += fmt.Sprintf(" LIMIT %s", limit)
	}
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var reviews []dto.Reviews
	for rows.Next() {
		var r dto.Reviews
		err := rows.Scan(&r.Id, &r.Author, &r.Date, &r.Text, &r.Photo)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, r)
	}
	return reviews, nil
}
