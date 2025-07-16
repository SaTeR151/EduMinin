package restlk

import (
	"net/http"

	"github.com/SaTeR151/EduMinin/internal/apperror"
	"github.com/SaTeR151/EduMinin/internal/controller/rest/dto"
	"github.com/SaTeR151/EduMinin/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetLkInfo(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("getting lk info")
		accessTCokie, err := c.Request.Cookie("EduMininAT")
		if err != nil {
			logrus.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		lkInfo, err := services.LkManager.GetLkInfo(accessTCokie.Value)
		if err != nil {
			if err == apperror.ErrUnauthorized {
				logrus.Warn(err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error{Error: apperror.ErrUnauthorized.Error()})
				return
			}
			logrus.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, lkInfo)
		logrus.Info("lk info sent")
	}
}

func PostCourse(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("saving user course")
		accessTCokie, err := c.Request.Cookie("EduMininAT")
		if err != nil {
			logrus.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var userCourse dto.SaveUserCourse
		if err := c.ShouldBindJSON(&userCourse); err != nil {
			logrus.Warn(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIncorrectRequestBody.Error()})
			return
		}
		if userCourse.Id == "" {
			logrus.Warn(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIncorrectRequestBody.Error()})
			return
		}
		if err := services.LkManager.AddCourse(accessTCokie.Value, userCourse.Id); err != nil {
			if err == apperror.ErrUnauthorized {
				logrus.Warn(err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error{Error: apperror.ErrUnauthorized.Error()})
				return
			}
			logrus.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
			return
		}
		c.Status(http.StatusCreated)
		logrus.Info("user course saved")
	}
}

func PutEmail(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("changing email")
		accessTCokie, err := c.Request.Cookie("EduMininAT")
		if err != nil {
			logrus.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var changing dto.ChangingData
		if err := c.ShouldBindJSON(&changing); err != nil {
			logrus.Warn(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIncorrectRequestBody.Error()})
			return
		}
		if changing.Changing == "" {
			logrus.Warn(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIncorrectRequestBody.Error()})
			return
		}
		if err := services.LkManager.ChangeEmail(accessTCokie.Value, changing.Changing); err != nil {
			if err == apperror.ErrUnauthorized {
				logrus.Warn(err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error{Error: apperror.ErrUnauthorized.Error()})
				return
			}
			logrus.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
			return
		}
		c.Status(http.StatusOK)
		logrus.Info("user course saved")
	}
}

func PutPhone(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("changing email")
		accessTCokie, err := c.Request.Cookie("EduMininAT")
		if err != nil {
			logrus.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var changing dto.ChangingData
		if err := c.ShouldBindJSON(&changing); err != nil {
			logrus.Warn(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIncorrectRequestBody.Error()})
			return
		}
		if changing.Changing == "" {
			logrus.Warn(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIncorrectRequestBody.Error()})
			return
		}
		if err := services.LkManager.ChangePhone(accessTCokie.Value, changing.Changing); err != nil {
			if err == apperror.ErrUnauthorized {
				logrus.Warn(err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error{Error: apperror.ErrUnauthorized.Error()})
				return
			}
			logrus.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
			return
		}
		c.Status(http.StatusOK)
		logrus.Info("user course saved")
	}
}

func PutPhoto(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("changing email")
		accessTCokie, err := c.Request.Cookie("EduMininAT")
		if err != nil {
			logrus.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var changing dto.ChangingData
		if err := c.ShouldBindJSON(&changing); err != nil {
			logrus.Warn(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIncorrectRequestBody.Error()})
			return
		}
		if changing.Changing == "" {
			logrus.Warn(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIncorrectRequestBody.Error()})
			return
		}
		if err := services.LkManager.ChangePhoto(accessTCokie.Value, changing.Changing); err != nil {
			if err == apperror.ErrUnauthorized {
				logrus.Warn(err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error{Error: apperror.ErrUnauthorized.Error()})
				return
			}
			logrus.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
			return
		}
		c.Status(http.StatusOK)
		logrus.Info("user course saved")
	}
}

func PutFio(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("changing email")
		accessTCokie, err := c.Request.Cookie("EduMininAT")
		if err != nil {
			logrus.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var changing dto.ChangingData
		if err := c.ShouldBindJSON(&changing); err != nil {
			logrus.Warn(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIncorrectRequestBody.Error()})
			return
		}
		if changing.Changing == "" {
			logrus.Warn(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIncorrectRequestBody.Error()})
			return
		}
		if err := services.LkManager.ChangeFio(accessTCokie.Value, changing.Changing); err != nil {
			if err == apperror.ErrUnauthorized {
				logrus.Warn(err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error{Error: apperror.ErrUnauthorized.Error()})
				return
			}
			logrus.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
			return
		}
		c.Status(http.StatusOK)
		logrus.Info("user course saved")
	}
}
