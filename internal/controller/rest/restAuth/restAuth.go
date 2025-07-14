package restauth

import (
	"net/http"

	"github.com/SaTeR151/EduMinin/internal/apperror"
	"github.com/SaTeR151/EduMinin/internal/controller/rest/dto"
	"github.com/SaTeR151/EduMinin/internal/controller/rest/restutils"
	"github.com/SaTeR151/EduMinin/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SignUp(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("getting tokens")
		var userData dto.UserData
		if err := c.ShouldBindJSON(&userData); err != nil {
			logrus.Warn(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIncorrectRequestBody.Error()})
			return
		}
		aToken, rToken, err := services.AuthManager.Signup(userData)
		if err != nil {
			if err == apperror.ErrUnauthorized {
				logrus.Warn(err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error{Error: err.Error()})
				return
			}
			logrus.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
			return
		}

		restutils.SetCookieTokens(c, aToken, rToken)
		c.Status(http.StatusCreated)
		logrus.Info("tokens have been sent")
	}
}

func Logout(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("starting logout")
		atCookie, err := c.Request.Cookie("at")
		if err != nil {
			logrus.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error{Error: apperror.ErrUnauthorized.Error()})
			return
		}

		err = services.AuthManager.Logout(atCookie.Value)
		if err != nil {
			logrus.Warn(err)
			if err != apperror.ErrUnauthorized {
				c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
				return
			}
		}
		c.Status(http.StatusNoContent)
		logrus.Info("logout finished")
	}
}

func Register(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("getting tokens")
		var userData dto.UserData
		if err := c.ShouldBindJSON(&userData); err != nil {
			logrus.Warn(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{Error: apperror.ErrIncorrectRequestBody.Error()})
			return
		}
		err := services.AuthManager.Register(userData)
		if err != nil {
			if err == apperror.ErrUnauthorized {
				logrus.Warn(err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error{Error: err.Error()})
				return
			}
			logrus.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
			return
		}
		c.Status(http.StatusCreated)
		logrus.Info("tokens have been sent")
	}
}
