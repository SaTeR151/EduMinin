package authmanager

import (
	"github.com/SaTeR151/EduMinin/internal/controller/rest/dto"
	"github.com/SaTeR151/EduMinin/internal/database/sqlite"
)

type AuthManagerService interface {
	Register(data dto.UserData) (err error)
	Signup(data dto.UserData) (aToken string, rToken string, err error)
	Logout(aToken string) (err error)
}

type AuthManager struct {
	db sqlite.SQLiteManagerDB
}

func New(db sqlite.SQLiteManagerDB) *AuthManager {
	return &AuthManager{db: db}
}

func (am *AuthManager) Register(data dto.UserData) (err error) {
	return nil
}

func (am *AuthManager) Signup(data dto.UserData) (aToken string, rToken string, err error) {
	return aToken, rToken, err
}

func (am *AuthManager) Logout(aToken string) (err error) {
	return err
}
