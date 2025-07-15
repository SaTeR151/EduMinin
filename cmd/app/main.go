package main

import (
	"github.com/SaTeR151/EduMinin/internal/config"
	restauth "github.com/SaTeR151/EduMinin/internal/controller/rest/restAuth"
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

	// router.Static("/main/assets", "./web/assets") // если они у вас есть
	// router.Static("/main/css", "./web/css")       // если они у вас есть
	// router.Static("/main/js", "./web/js")         // если они у вас есть

	// // 2. Правило «/main/<page> -> ./web/<page>.html»
	// router.GET("/main/:page", func(c *gin.Context) {
	// 	page := c.Param("page")
	// 	// защита от ".." и лишних слэшей
	// 	safe := filepath.Clean(page)
	// 	if strings.ContainsRune(safe, os.PathSeparator) {
	// 		c.AbortWithStatus(http.StatusNotFound)
	// 		return
	// 	}

	// 	file := filepath.Join("web", safe+".html")
	// 	if _, err := os.Stat(file); err == nil {
	// 		c.File(file)
	// 		return
	// 	}
	// 	c.AbortWithStatus(http.StatusNotFound)
	// })

	// // 3. Главная: /main/ (или просто /) → index.html
	// router.GET("/main/", func(c *gin.Context) {
	// 	c.File("./web/index.html")
	// })
	// router.GET("/", func(c *gin.Context) { // по желанию
	// 	c.Redirect(http.StatusMovedPermanently, "/main/")
	// })

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
		{
			authGroup := apiGroup.Group("/auth")
			authGroup.POST("/signup", restauth.Signup(services))
			authGroup.POST("/register", restauth.Register(services))
			authGroup.POST("/logout", restauth.Logout(services))
		}

	}
	if err := router.Run(":" + serverConfig.Port); err != nil {
		logrus.Error(err)
		return
	}
}
