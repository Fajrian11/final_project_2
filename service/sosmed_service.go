package service

import (
	"FP2/repositories"

	"github.com/gin-gonic/gin"
)

type SosmedService struct {
	rr repositories.SosmedRepoApi
}

func NewSosmedService(rr repositories.SosmedRepoApi) *SosmedService { //provie service
	return &SosmedService{rr: rr}
}

type SosmedServiceApi interface {
	CreateSosmedService(c *gin.Context) gin.H
	GetAllSosmedService(c *gin.Context) gin.H
	UpdateSosmedService(c *gin.Context) gin.H
	DeleteSosmedService(c *gin.Context) gin.H
}

func (ss SosmedService) CreateSosmedService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	Sosmed, _ := ss.rr.CreateSosmed(c)
	if Sosmed.Name == "" {
		result = gin.H{
			"error": "Your name is required",
		}
	} else if Sosmed.SocialMediaURL == "" {
		result = gin.H{
			"error": "Your social_media_url is required",
		}
	} else {
		result = gin.H{
			"Success":          "Data Has been created",
			"id":               Sosmed.ID,
			"name":             Sosmed.Name,
			"social_media_url": Sosmed.SocialMediaURL,
			"user_id":          Sosmed.UserID,
			"created_at":       Sosmed.CreatedAt,
		}
	}
	return result
}

func (ss SosmedService) GetAllSosmedService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	GetAllSosmed, err := ss.rr.GetAllSosmed(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"result": GetAllSosmed,
			"count":  len(GetAllSosmed),
		}
	}
	return result
}

func (ss SosmedService) UpdateSosmedService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	Sosmed, err := ss.rr.UpdateSosmed(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"Success":          "Data Has been Updated",
			"id":               Sosmed.ID,
			"name":             Sosmed.Name,
			"social_media_url": Sosmed.SocialMediaURL,
			"user_id":          Sosmed.UserID,
			"updated_at":       Sosmed.UpdatedAt,
		}
	}
	return result
}

func (ss SosmedService) DeleteSosmedService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	_, err := ss.rr.DeleteSosmed(c)
	if err != nil {
		result = gin.H{
			"result":  "Gagal Menghapus Data",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"Success": "Your Social Media has been successfully deleted",
		}
	}
	return result
}
