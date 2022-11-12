package service

import (
	"FP2/repositories"

	"github.com/gin-gonic/gin"
)

type PhotoService struct {
	rr repositories.PhotoRepoApi
}

func NewPhotoService(rr repositories.PhotoRepoApi) *PhotoService { //provie service
	return &PhotoService{rr: rr}
}

type PhotoServiceApi interface {
	CreatePhotoService(c *gin.Context) gin.H
	GetAllPhotoService(c *gin.Context) gin.H
	UpdatePhotoService(c *gin.Context) gin.H
	DeletePhotoService(c *gin.Context) gin.H
}

func (ps PhotoService) CreatePhotoService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	Photo, _ := ps.rr.CreatePhoto(c)
	if Photo.Title == "" {
		result = gin.H{
			"error": "Your Title is required",
		}
	} else if Photo.Caption == "" {
		result = gin.H{
			"error": "Your Caption is required",
		}
	} else if Photo.PhotoURL == "" {
		result = gin.H{
			"error": "Your PhotoURL is required",
		}
	} else {
		result = gin.H{
			"Success":    "Data Has been created",
			"id":         Photo.ID,
			"title":      Photo.Title,
			"caption":    Photo.Caption,
			"photo_url":  Photo.PhotoURL,
			"user_id":    Photo.UserID,
			"created_at": Photo.CreatedAt,
		}
	}
	return result
}

func (ps PhotoService) GetAllPhotoService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	GetAllPhoto, err := ps.rr.GetAllPhoto(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"result": GetAllPhoto,
			"count":  len(GetAllPhoto),
		}
	}
	return result
}

func (ps PhotoService) UpdatePhotoService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	Photo, err := ps.rr.UpdatePhoto(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"Success":    "Data Has been Updated",
			"id":         Photo.ID,
			"title":      Photo.Title,
			"caption":    Photo.Caption,
			"photo_url":  Photo.PhotoURL,
			"user_id":    Photo.UserID,
			"created_at": Photo.CreatedAt,
		}
	}
	return result
}

func (ps PhotoService) DeletePhotoService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	_, err := ps.rr.DeletePhoto(c)
	if err != nil {
		result = gin.H{
			"result":  "Gagal Menghapus Data",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"Success": "Your Photo has been successfully deleted",
		}
	}
	return result
}
