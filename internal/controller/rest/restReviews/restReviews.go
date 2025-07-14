package restreviews

import (
	"net/http"

	"github.com/SaTeR151/EduMinin/internal/apperror"
	"github.com/SaTeR151/EduMinin/internal/controller/rest/dto"
	"github.com/SaTeR151/EduMinin/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func PostReview(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("saving review")
		var review dto.Reviews
		if err := c.ShouldBindJSON(&review); err != nil {
			logrus.Warn(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIncorrectRequestBody.Error()})
			return
		}
		if err := services.ReviewsManager.PostReview(review); err != nil {
			logrus.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
			return
		}
		c.Status(http.StatusCreated)
		logrus.Info("review saved")
	}
}

func GetReviews(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("getting reviews")
		limit := c.Query("limit")
		data, err := services.ReviewsManager.GetReviews(limit)
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
		logrus.Info("reviews sent")
	}
}

func DeleteReview(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("deleting review")
		id := c.Query("id")
		if id == "" {
			logrus.Warn(apperror.ErrIdRequired)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIdRequired.Error()})
			return
		}
		err := services.ReviewsManager.DeleteReview(id)
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
		logrus.Info("review deleted")
	}
}
