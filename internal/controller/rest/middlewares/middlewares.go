package middlewares

import (
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/SaTeR151/EduMinin/internal/apperror"
	"github.com/SaTeR151/EduMinin/internal/controller/rest/restutils"
	"github.com/SaTeR151/EduMinin/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func CheckAuthorization(services *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("checking authorization")
		rtCookie, err := c.Request.Cookie("EduMininRT")
		if err != nil {
			logrus.Error(err)
			//c.AbortWithStatus(http.StatusUnauthorized)
			c.Redirect(http.StatusFound, "/main/login") // +
			c.Abort()                                   // +
			return
		}
		rTokenCookie, err := url.QueryUnescape(rtCookie.Value)
		if err != nil {
			logrus.Error(err)
			//c.AbortWithStatus(http.StatusBadRequest)
			c.Redirect(http.StatusFound, "/main/login") // +
			c.Abort()                                   // +
			return
		}
		gettingRTBase64, err := base64.StdEncoding.DecodeString(rTokenCookie)
		if err != nil {
			logrus.Error(err)
			//c.AbortWithStatus(http.StatusBadRequest)
			c.Redirect(http.StatusFound, "/main/login") // +
			c.Abort()                                   // +
			return
		}

		accessTCokie, err := c.Request.Cookie("EduMininAT")
		if err != nil {
			logrus.Error(err)
			//c.AbortWithStatus(http.StatusUnauthorized)
			c.Redirect(http.StatusFound, "/main/login") // +
			c.Abort()                                   // +
			return
		}

		if err = services.AuthManager.CheckTokens(accessTCokie.Value, string(gettingRTBase64)); err != nil {
			if err != jwt.ErrTokenExpired {
				logrus.Error(err)
				//c.AbortWithStatus(http.StatusUnauthorized)
				c.Redirect(http.StatusFound, "/main/login") // +
				c.Abort()                                   // +
				return
			}
			logrus.Info("access token expired")
			aToken, rToken, err := services.AuthManager.RefreshTokens(accessTCokie.Value, string(gettingRTBase64), c.Request.Header.Get("User-Agent"), c.ClientIP())
			if err != nil {
				logrus.Warn(err)
				if err == apperror.ErrUnauthorized {
					//c.AbortWithStatus(http.StatusUnauthorized)
					c.Redirect(http.StatusFound, "/main/login") // +
					c.Abort()                                   // +
					return
				}
				//c.AbortWithStatus(http.StatusInternalServerError)
				c.Redirect(http.StatusFound, "/main/login") // +
				c.Abort()                                   // +
				return
			}
			restutils.SetCookieTokens(c, aToken, rToken)
			c.Next()
			return
		}
		c.Next()
		return
	}
}
