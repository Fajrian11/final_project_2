package models

import (
	"FP2/helpers"
	"errors"

	"github.com/asaskevich/govalidator"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null,index:username,unique" json:"username" form:"username" valid:"required~Your Username is required"`
	Email    string `gorm:"not null" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password minimal harus 6 karakter"`
	Age      uint   `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required"`
}

// validasi field field di database
func (u *User) BeforeCreate() (err error) {
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

func (u *User) BeforeUpdate() (err error) {
	_, errUpdate := govalidator.ValidateStruct(u)

	if errUpdate != nil {
		err = errUpdate
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
