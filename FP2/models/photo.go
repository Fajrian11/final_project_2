package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID       uint
	Title    string
	Caption  string
	PhotoURL string
	UserID   uint
	User     *User
	TimeModel
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)
	if errUpdate != nil {
		err = errUpdate
		return
	}
	err = nil
	return
}
