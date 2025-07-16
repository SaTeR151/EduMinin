package authmanager

import (
	"crypto/sha256"
	"database/sql"
	"fmt"

	"github.com/SaTeR151/EduMinin/internal/apperror"
	"github.com/SaTeR151/EduMinin/internal/controller/rest/dto"
	"github.com/SaTeR151/EduMinin/internal/database/sqlite"
	"github.com/SaTeR151/EduMinin/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type AuthManagerService interface {
	Register(data dto.UserData) (err error)
	Signup(data dto.UserData, userAgent string, ip string) (aToken string, rToken string, err error)
	Logout(aToken string) (err error)
	RefreshTokens(at string, rt string, userAgent string, ip string) (aToken string, rToken string, err error)
	CheckTokens(aToken string, rToken string) error
}

type AuthManager struct {
	db sqlite.SQLiteManagerDB
}

func New(db sqlite.SQLiteManagerDB) *AuthManager {
	return &AuthManager{db: db}
}

func (am *AuthManager) Register(data dto.UserData) (err error) {
	err = am.db.CheckUsersLoginExists(fmt.Sprintf("%x", sha256.Sum256([]byte(data.Login))))
	if err != nil && err != apperror.ErrUncorrectData {
		return err
	}
	return am.db.RegisterUser(fmt.Sprintf("%x", sha256.Sum256([]byte(data.Login))), fmt.Sprintf("%x", sha256.Sum256([]byte(data.Pass))), data.Login)
}

func (am *AuthManager) Signup(data dto.UserData, userAgent string, ip string) (aToken string, rToken string, err error) {
	login := fmt.Sprintf("%x", sha256.Sum256([]byte(data.Login)))
	logrus.Debug("checking user exists")
	err = am.db.CheckUsersLoginExists(login)
	if err != nil {
		return aToken, rToken, err
	}
	logrus.Debug("getting user info")
	oldLogin, oldPass, err := am.db.GetUserInfo(login)
	if err != nil {
		return aToken, rToken, err
	}
	if oldLogin != login || oldPass != fmt.Sprintf("%x", sha256.Sum256([]byte(data.Pass))) {
		return aToken, rToken, apperror.ErrUserNotFound

	}
	logrus.Debug("checking user authorization")

	err = am.db.CheckUserAuthorization(login)
	if err != nil && err != apperror.ErrUserAlreadyAuthorization {
		return aToken, rToken, err
	}
	if err == apperror.ErrUserAlreadyAuthorization {
		err = am.db.DeleteAuthorizationUser(login)
	}
	logrus.Debug("generating tokens")

	aToken, rToken, err = utils.NewTokens(userAgent, data.Login)
	if err != nil {
		return aToken, rToken, err
	}
	logrus.Debug("saveing auth user")

	err = am.db.Signup(login, fmt.Sprintf("%x", sha256.Sum256([]byte(rToken))))
	if err != nil {
		if err == sql.ErrNoRows {
			return aToken, rToken, apperror.ErrUnauthorized

		}
		return aToken, rToken, err
	}
	return aToken, rToken, err
}

func (am *AuthManager) Logout(aToken string) (err error) {
	login, err := utils.GetLoginFromJWT(aToken)
	if err != nil {
		return err
	}
	return am.db.DeleteAuthorizationUser(fmt.Sprintf("%x", sha256.Sum256([]byte(login))))
}

func (am *AuthManager) RefreshTokens(at string, rt string, userAgent string, ip string) (aToken string, rToken string, err error) {
	logrus.Debug("refreshing tokens")

	login, err := utils.GetLoginFromJWT(at)
	if err != nil {
		return aToken, rToken, err
	}

	ok, err := am.CompareRT(rt, login)
	if err != nil {
		return aToken, rToken, err
	}
	if !ok {
		return aToken, rToken, apperror.ErrUnauthorized
	}

	logrus.Debug("generating new tokens")
	aToken, rToken, err = utils.NewTokens(userAgent, login)
	if err != nil {
		return aToken, rToken, err
	}

	err = am.db.Signup(fmt.Sprintf("%x", sha256.Sum256([]byte(login))), fmt.Sprintf("%x", sha256.Sum256([]byte(rToken))))
	if err != nil {
		return aToken, rToken, err
	}

	logrus.Debug("tokens refreshed")
	return aToken, rToken, nil
}

func (am *AuthManager) CompareRT(rToken, login string) (bool, error) {
	rToken = fmt.Sprintf("%x", sha256.Sum256([]byte(rToken)))
	rTokensOld, err := am.db.GetTokens(fmt.Sprintf("%x", sha256.Sum256([]byte(login))))
	if err != nil {
		if err == sql.ErrNoRows {
			return false, apperror.ErrUnauthorized
		}
		return false, err
	}
	for _, v := range rTokensOld {
		if rToken == v {
			return true, nil
		}
	}
	return false, nil
}

func (am *AuthManager) CheckTokens(aToken string, rToken string) error {
	err := utils.CheckLinkTokens(aToken, rToken)
	if err != nil {
		if err.Error() == "token has invalid claims: token is expired" {
			return jwt.ErrTokenExpired
		}
		return err
	}
	return nil
}
