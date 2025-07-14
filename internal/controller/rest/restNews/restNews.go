package restnews

import (
	"net/http"

	"github.com/SaTeR151/EduMinin/internal/apperror"
	"github.com/SaTeR151/EduMinin/internal/controller/rest/dto"
	"github.com/SaTeR151/EduMinin/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func PostNews(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("saving news")
		var news dto.News
		if err := c.ShouldBindJSON(&news); err != nil {
			logrus.Warn(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIncorrectRequestBody.Error()})
			return
		}
		if err := services.NewsManager.PostNews(news); err != nil {
			logrus.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
			return
		}
		c.Status(http.StatusCreated)
		logrus.Info("news saved")
	}
}

func GetNews(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("getting news")
		limit := c.Query("limit")
		data, err := services.NewsManager.GetNews(limit)
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
		logrus.Info("news sent")
	}
}

func DeleteNews(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("deleting news")
		id := c.Query("id")
		if id == "" {
			logrus.Warn(apperror.ErrIdRequired)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIdRequired.Error()})
			return
		}
		err := services.NewsManager.DeleteNews(id)
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
		logrus.Info("news deleted")
	}
}
