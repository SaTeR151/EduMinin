package lkmanager

import (
	"crypto/sha256"
	"fmt"

	"github.com/SaTeR151/EduMinin/internal/apperror"
	"github.com/SaTeR151/EduMinin/internal/controller/rest/dto"
	"github.com/SaTeR151/EduMinin/internal/database/sqlite"
	"github.com/SaTeR151/EduMinin/internal/utils"
	"github.com/sirupsen/logrus"
)

type LkManagerService interface {
	GetLkInfo(aToken string) (dto.LK, error)
	ChangePhoto(aToken string, photo string) error
	ChangeFio(aToken string, fio string) error
	ChangePhone(aToken string, phone string) error
	ChangeEmail(aToken string, email string) error
	AddCourse(aToken string, id string) error
}

type LkManager struct {
	db sqlite.SQLiteManagerDB
}

func New(db sqlite.SQLiteManagerDB) *LkManager {
	return &LkManager{db: db}
}

func (lkm *LkManager) GetLkInfo(aToken string) (dto.LK, error) {
	var lk dto.LK
	login, err := utils.GetLoginFromJWT(aToken)
	if err != nil {
		return lk, apperror.ErrUnauthorized
	}
	login = fmt.Sprintf("%x", sha256.Sum256([]byte(login)))
	lk, err = lkm.db.GetLKInfo(login)
	if err != nil {
		return lk, err
	}
	logrus.Debug("gettig user courses")
	lk.Courses, err = lkm.db.GetUserCourses(login)
	if err != nil {
		return lk, err
	}
	return lk, nil
}

func (lkm *LkManager) ChangePhoto(aToken string, photo string) error {
	login, err := utils.GetLoginFromJWT(aToken)
	if err != nil {
		return apperror.ErrUnauthorized
	}
	login = fmt.Sprintf("%x", sha256.Sum256([]byte(login)))
	return lkm.db.ChangePhoto(login, photo)
}

func (lkm *LkManager) ChangeFio(aToken string, fio string) error {
	login, err := utils.GetLoginFromJWT(aToken)
	if err != nil {
		return apperror.ErrUnauthorized
	}
	login = fmt.Sprintf("%x", sha256.Sum256([]byte(login)))
	return lkm.db.ChangeFIO(login, fio)
}

func (lkm *LkManager) ChangePhone(aToken string, phone string) error {
	login, err := utils.GetLoginFromJWT(aToken)
	if err != nil {
		return apperror.ErrUnauthorized
	}
	login = fmt.Sprintf("%x", sha256.Sum256([]byte(login)))
	return lkm.db.ChangePhone(login, phone)
}
func (lkm *LkManager) ChangeEmail(aToken string, email string) error {
	login, err := utils.GetLoginFromJWT(aToken)
	if err != nil {
		return apperror.ErrUnauthorized
	}
	login = fmt.Sprintf("%x", sha256.Sum256([]byte(login)))
	return lkm.db.ChangeEmail(login, email)
}

func (lkm *LkManager) AddCourse(aToken string, id string) error {
	login, err := utils.GetLoginFromJWT(aToken)
	if err != nil {
		return apperror.ErrUnauthorized
	}
	login = fmt.Sprintf("%x", sha256.Sum256([]byte(login)))
	return lkm.db.AddUserCourse(login, id)
}
