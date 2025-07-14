package restreviews

import (
	"github.com/SaTeR151/EduMinin/internal/services"
	"github.com/gin-gonic/gin"
)

func PostReview(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		// logrus.Info("saving review")
		// var review dto.Course
		// if err := c.ShouldBindJSON(&review); err != nil {
		// 	logrus.Warn(err)
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIncorrectRequestBody.Error()})
		// 	return
		// }
		// if err := services.ReviewsManager.PostReview(review); err != nil {
		// 	logrus.Error(err)
		// 	c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
		// 	return
		// }
		// c.Status(http.StatusCreated)
		// logrus.Info("news review")
	}
}

func GetReviews(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func DeleteReview(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
