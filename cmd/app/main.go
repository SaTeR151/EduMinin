package main

import (
	"github.com/SaTeR151/EduMinin/internal/config"
	restcourses "github.com/SaTeR151/EduMinin/internal/controller/rest/restCourses"
	restevents "github.com/SaTeR151/EduMinin/internal/controller/rest/restEvents"
	restnews "github.com/SaTeR151/EduMinin/internal/controller/rest/restNews"
	restreviews "github.com/SaTeR151/EduMinin/internal/controller/rest/restReviews"
	"github.com/SaTeR151/EduMinin/internal/database/sqlite"
	"github.com/SaTeR151/EduMinin/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load()
	config.InitLoggerConfig()
	logrus.Info("starting server")
	sqliteConfig := config.GetSQLiteConfig()
	db, close, err := sqlite.Open(sqliteConfig)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer close()
	err = db.MigrationUP()
	if err != nil {
		logrus.Error(err)
		return
	}
	services := services.New(db)
	serverConfig := config.GetServerConfig()

	router := gin.Default()

	router.Static("/main", "./web")

	{
		apiGroup := router.Group("/api")
		{
			apiGroup.GET("/news", restnews.GetNews(services))
			newsGroups := apiGroup.Group("/news")
			newsGroups.POST("/", restnews.PostNews(services))
			newsGroups.DELETE("/", restnews.DeleteNews(services))
		}
		{
			apiGroup.GET("/courses", restcourses.GetCourses(services))
			apiGroup.GET("/courses/title", restcourses.GetCoursesTitle(services))
			coursesGroups := apiGroup.Group("/courses")
			coursesGroups.POST("/", restcourses.PostCourse(services))
			coursesGroups.DELETE("/", restcourses.DeleteCourse(services))
		}
		{
			apiGroup.GET("/reviews", restreviews.GetReviews(services))
			reviewsGroups := apiGroup.Group("/reviews")
			reviewsGroups.POST("/", restreviews.PostReview(services))
			reviewsGroups.DELETE("/", restreviews.DeleteReview(services))
		}
		{
			apiGroup.GET("/events", restevents.GetEvents(services))
			eventsGroups := apiGroup.Group("/events")
			eventsGroups.POST("/", restevents.PostEvent(services))
			eventsGroups.DELETE("/", restevents.DeleteEvent(services))
		}

	}
	if err := router.Run(":" + serverConfig.Port); err != nil {
		logrus.Error(err)
		return
	}
}
