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
	GetCoursesTitle() ([]dto.CoursesTitle, error)

	SaveNews(news dto.News) error
	DeleteNews(id string) error
	GetNews(limit string) ([]dto.News, error)

	SaveReview(review dto.Reviews) error
	DeleteReview(id string) error
	GetReviews(limit string) ([]dto.Reviews, error)

	SaveEvent(events dto.Events) error
	DeleteEvent(id string) error
	GetEvents(limit string) ([]dto.Events, error)

	CheckUsersLoginExists(login string) error
	RegisterUser(login string, pass string) error

	CheckUserAuthorization(login string) error
	DeleteAuthorizationUser(login string) error
	GetUserInfo(login string) (oldLogin string, oldPass string, err error)
	Signup(login string, rt string) error
	GetTokens(login string) (rTokens []string, err error)

	ChangeEmail(login string, email string) error
	ChangeFIO(login string, fio string) error
	ChangePhone(login string, phone string) error
	ChangePhoto(login string, photo string) error
	AddUserCourse(login string, id string) error
	GetUserCourses(login string) ([]dto.Course, error)
	GetLKInfo(login string) (dto.LK, error)
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
	var courses []dto.Course
	for rows.Next() {
		var c dto.Course
		err := rows.Scan(&c.Id, &c.Title, &c.Description, &c.AcademicHours, &c.Price, &c.Image)
		if err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}
	return courses, nil
}

func (db *SQLiteManager) GetCoursesTitle() ([]dto.CoursesTitle, error) {
	rows, err := db.db.Query("SELECT title FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var coursesTitle []dto.CoursesTitle
	for rows.Next() {
		var c dto.CoursesTitle
		err := rows.Scan(&c.Title)
		if err != nil {
			return nil, err
		}
		coursesTitle = append(coursesTitle, c)
	}
	return coursesTitle, nil
}

func (db *SQLiteManager) SaveNews(news dto.News) error {
	_, err := db.db.Exec("INSERT INTO news (title, description, content, date, image_path) values (:title, :description, :content, :date, :image_path)",
		sql.Named("title", news.Title),
		sql.Named("description", news.Description),
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
	query := "SELECT id, title, description, content, date, image_path image FROM news"
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
		var c dto.News
		err := rows.Scan(&c.Id, &c.Title, &c.Description, &c.Text, &c.Date, &c.Image)
		if err != nil {
			return nil, err
		}
		news = append(news, c)
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
	query := "SELECT id, author_name, date, review_text, photo_path FROM reviews"
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

func (db *SQLiteManager) SaveEvent(events dto.Events) error {
	_, err := db.db.Exec("INSERT INTO events (title, content, date, image_path) values (:title, :content, :date, :image_path)",
		sql.Named("title", events.Title),
		sql.Named("content", events.Text),
		sql.Named("date", events.Date),
		sql.Named("image_path", events.Image),
	)
	return err
}

func (db *SQLiteManager) DeleteEvent(id string) error {
	res, err := db.db.Exec("DELETE FROM events WHERE id = :id", sql.Named("id", id))
	if err != nil {
		return nil
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return nil
	}
	if rows == 0 {
		return apperror.ErrDataNotFound
	}
	return err
}

func (db *SQLiteManager) GetEvents(limit string) ([]dto.Events, error) {
	query := "SELECT id, title, content, date, image_path FROM events"
	if limit != "" {
		query += fmt.Sprintf(" LIMIT %s", limit)
	}
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []dto.Events
	for rows.Next() {
		var c dto.Events
		err := rows.Scan(&c.Id, &c.Title, &c.Text, &c.Date, &c.Image)
		if err != nil {
			return nil, err
		}
		events = append(events, c)
	}
	return events, nil
}

func (db *SQLiteManager) CheckUsersLoginExists(login string) error {
	rows, err := db.db.Query("SELECT count(login) count FROM users WHERE login = :login", sql.Named("login", login))
	if err != nil {
		return err
	}
	defer rows.Close()
	var count int
	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return nil
		}
	}
	if count == 0 {
		return apperror.ErrUncorrectData
	}
	return nil
}

func (db *SQLiteManager) RegisterUser(login string, pass string) error {
	_, err := db.db.Exec("INSERT INTO users (login, pass, fio) values (:login, :pass, :login)",
		sql.Named("login", login),
		sql.Named("pass", pass),
	)
	if err != nil {
		return apperror.ErrUserAlreadyExists
	}
	return nil
}

func (db *SQLiteManager) CheckUserAuthorization(login string) error {
	res, err := db.db.Exec("SELECT login FROM auth WHERE login = :login", sql.Named("login", login))
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 0 {
		return apperror.ErrUserAlreadyAuthorization
	}
	return nil
}

func (db *SQLiteManager) DeleteAuthorizationUser(login string) error {
	res, err := db.db.Exec("DELETE FROM auth WHERE login = :login", sql.Named("login", login))
	if err != nil {
		return nil
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return nil
	}
	if rows == 0 {
		return apperror.ErrUserNotFound
	}
	return err
}

func (db *SQLiteManager) GetUserInfo(login string) (oldLogin string, oldPass string, err error) {
	rows, err := db.db.Query("SELECT login, pass FROM users WHERE login = :login", sql.Named("login", login))
	if err != nil {
		return oldLogin, oldPass, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&oldLogin, &oldPass)
		if err != nil {
			return oldLogin, oldPass, err

		}
	}
	return oldLogin, oldPass, err
}

func (db *SQLiteManager) Signup(login string, rt string) error {
	_, err := db.db.Exec("INSERT INTO auth (login, refresh_token) values (:login, :refresh_token)",
		sql.Named("login", login),
		sql.Named("refresh_token", rt),
	)
	return err
}

func (db *SQLiteManager) GetTokens(login string) (rTokens []string, err error) {
	rows, err := db.db.Query("SELECT refresh_token FROM auth WHERE login = :login", sql.Named("login", login))
	if err != nil {
		return rTokens, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&rTokens)
		if err != nil {
			return rTokens, err

		}
	}
	return rTokens, err
}

func (db *SQLiteManager) ChangeEmail(login string, email string) error {
	_, err := db.db.Exec("UPDATE users SET email = :email WHERE login = :login",
		sql.Named("login", login),
		sql.Named("email", email),
	)
	return err
}

func (db *SQLiteManager) ChangeFIO(login string, fio string) error {
	_, err := db.db.Exec("UPDATE users SET fio = :fio WHERE login = :login",
		sql.Named("login", login),
		sql.Named("fio", fio),
	)
	return err
}
func (db *SQLiteManager) ChangePhone(login string, phone string) error {
	_, err := db.db.Exec("UPDATE users SET phone = :phone WHERE login = :login",
		sql.Named("login", login),
		sql.Named("phone", phone),
	)
	return err
}
func (db *SQLiteManager) ChangePhoto(login string, photo string) error {
	_, err := db.db.Exec("UPDATE users SET photo = :photo WHERE login = :login",
		sql.Named("login", login),
		sql.Named("photo", photo),
	)
	return err
}

func (db *SQLiteManager) AddUserCourse(login string, id string) error {
	_, err := db.db.Exec("INSERT INTO users_courses (login, course_id) value (:login, :id)",
		sql.Named("login", login),
		sql.Named("id", id),
	)
	return err
}

func (db *SQLiteManager) GetUserCourses(login string) ([]dto.Course, error) {
	rows, err := db.db.Query("SELECT c.* FROM courses c JOIN users_courses uc ON c.id = uc.course_id WHERE uc.login = :login", sql.Named("login", login))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var courses []dto.Course
	for rows.Next() {
		var c dto.Course
		err := rows.Scan(&c.Id, &c.Title, &c.Description, &c.AcademicHours, &c.Price, &c.Image)
		if err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}
	return courses, nil
}

func (db *SQLiteManager) GetLKInfo(login string) (dto.LK, error) {
	var lk dto.LK
	rows, err := db.db.Query("SELECT fio, photo, email, phone WHERE login = :login", sql.Named("login", login))
	if err != nil {
		return lk, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&lk.Fio, &lk.Photo, &lk.Email, &lk.Phone)
		if err != nil {
			return lk, err
		}
	}
	return lk, nil
}
