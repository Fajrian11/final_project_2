package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name           string `json:"name" form:"name" validation:"required"`
	SocialMediaURL string `json:"social_media_url" form:"social_media_url" validation:"required"`
	UserID         uint
	User           *User
}

func (p *SocialMedia) BeforeCreate() (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *SocialMedia) BeforeUpdate() (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)
	if errUpdate != nil {
		err = errUpdate
		return
	}
	err = nil
	return
}
