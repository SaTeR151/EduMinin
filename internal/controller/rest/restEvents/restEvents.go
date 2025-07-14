package restevents

import (
	"net/http"

	"github.com/SaTeR151/EduMinin/internal/apperror"
	"github.com/SaTeR151/EduMinin/internal/controller/rest/dto"
	"github.com/SaTeR151/EduMinin/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func PostEvent(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("saving event")
		var event dto.Events
		if err := c.ShouldBindJSON(&event); err != nil {
			logrus.Warn(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIncorrectRequestBody.Error()})
			return
		}
		if err := services.EventsManager.PostEvent(event); err != nil {
			logrus.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
			return
		}
		c.Status(http.StatusCreated)
		logrus.Info("event saved")
	}
}

func GetEvents(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("getting events")
		limit := c.Query("limit")
		data, err := services.EventsManager.GetEvents(limit)
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
		logrus.Info("events sent")
	}
}

func DeleteEvent(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("deleting events")
		id := c.Query("id")
		if id == "" {
			logrus.Warn(apperror.ErrIdRequired)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIdRequired.Error()})
			return
		}
		err := services.EventsManager.DeleteEvent(id)
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
		logrus.Info("events deleted")
	}
}
