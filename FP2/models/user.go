package models

import (
	"FP2/helpers"
	"errors"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID       uint
	Username string
	Email    string
	Password string
	Age      uint
	TimeModel
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	if u.Age < 8 {
		err = errors.New("Minimum Age to register is 8")
		return err
	}
	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
