package restcourses

import (
	"net/http"

	"github.com/SaTeR151/EduMinin/internal/apperror"
	"github.com/SaTeR151/EduMinin/internal/controller/rest/dto"
	"github.com/SaTeR151/EduMinin/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func PostCourse(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("saving cours")
		var course dto.Course
		if err := c.ShouldBindJSON(&course); err != nil {
			logrus.Warn(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIncorrectRequestBody.Error()})
			return
		}
		if err := services.CoursesManager.PostCourse(course); err != nil {
			logrus.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
			return
		}
		c.Status(http.StatusCreated)
		logrus.Info("course saved")
	}
}

func GetCourses(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("getting courses")
		limit := c.Query("limit")
		data, err := services.CoursesManager.GetCourses(limit)
		if err != nil {
			logrus.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
			return
		}
		if len(data) == 0 {
			logrus.Warn(apperror.ErrDataNotFound)
			c.AbortWithStatusJSON(http.StatusNotFound, dto.Error{Error: apperror.ErrDataNotFound.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, dto.Data{Data: data})
		logrus.Info("courses sent")
	}
}

func DeleteCourse(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("deleting course")
		id := c.Query("id")
		if id == "" {
			logrus.Warn(apperror.ErrIdRequired)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIdRequired.Error()})
			return
		}
		err := services.CoursesManager.DeleteCourse(id)
		if err != nil {
			if err == apperror.ErrDataNotFound {
				logrus.Warn(apperror.ErrDataNotFound)
				c.AbortWithStatusJSON(http.StatusNoContent, dto.Error{Error: apperror.ErrDataNotFound.Error()})
				return
			}
			logrus.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
		logrus.Info("course deleted")
	}
}
