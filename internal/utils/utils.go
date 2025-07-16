package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/SaTeR151/EduMinin/internal/apperror"
	"github.com/golang-jwt/jwt/v5"
)

const str = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

func CreateLink() (string, error) {
	tokenLink := make([]byte, 32)
	_, err := rand.Read(tokenLink)
	if err != nil {
		return "", err
	}
	for i, j := range tokenLink {
		tokenLink[i] = str[j%byte(len(str))]
	}
	return string(tokenLink), nil
}

func NewTokens(userAgent string, login string) (aToken string, rToken string, err error) {
	tokenLink, err := CreateLink()
	if err != nil {
		return aToken, rToken, err
	}

	// create access token
	atTimeExp, err := strconv.Atoi(os.Getenv("ATEXPIRES"))
	if err != nil {
		return aToken, rToken, err
	}
	claims := &jwt.MapClaims{
		"exp":        time.Now().Add(time.Second * time.Duration(atTimeExp)).Unix(),
		"userAgent":  userAgent,
		"login":      login,
		"linkString": tokenLink,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	aToken, err = accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return aToken, rToken, err
	}

	// create refresh token
	rtSH := sha256.Sum256([]byte(tokenLink))
	rToken = fmt.Sprintf("%x%v", rtSH, aToken[len(aToken)-6:])
	return aToken, rToken, nil
}

func GetLoginFromJWT(aToken string) (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(aToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil && err.Error() != "token has invalid claims: token is expired" {
		return "", err
	}
	return claims["login"].(string), nil
}

func CheckLinkTokens(aToken string, rToken string) error {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(aToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}
	rToken = rToken[:len(rToken)-6]
	linkString := claims["linkString"].(string)
	if fmt.Sprintf("%x", sha256.Sum256([]byte(linkString))) != rToken {
		return apperror.ErrUnauthorized
	}
	return nil
}
