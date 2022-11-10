package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `gorm:"not null" json:"title" form:"title" validation:"required"`
	Caption  string `gorm:"not null" json:"caption" form:"caption" validation:"required"`
	PhotoURL string `gorm:"not null" json:"photo_url" form:"photo_url" validation:"required"`
	UserID   uint
	User     User
}

func (p *Photo) BeforeCreate() (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *Photo) BeforeUpdate() (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)
	if errUpdate != nil {
		err = errUpdate
		return
	}
	err = nil
	return
}
