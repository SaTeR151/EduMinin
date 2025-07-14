package restutils

import (
	"encoding/base64"
	"net/http"
	"os"
	"strconv"

	"github.com/SaTeR151/EduMinin/internal/controller/rest/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetCookieTokens(c *gin.Context, accessT string, refreshT string) {
	rtB64 := base64.StdEncoding.EncodeToString([]byte(refreshT))
	timeExp, err := strconv.Atoi(os.Getenv("COOKIEEXPIRES"))
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
		return
	}
	c.SetCookie("EduMininAT", accessT, timeExp, "/", "localhost", false, true)
	c.SetCookie("EduMininRT", rtB64, timeExp, "/", "localhost", true, true)
}
