package repositories

import (
	"FP2/helpers"
	"FP2/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PhotoRepo struct {
	db *gorm.DB
}

func NewPhotoRepo(db *gorm.DB) PhotoRepo {
	return PhotoRepo{
		db: db,
	}
}

type PhotoRepoApi interface {
	CreatePhoto(c *gin.Context) (models.Photo, error)
	GetAllPhoto(c *gin.Context) ([]models.Photo, error)
	UpdatePhoto(c *gin.Context) (models.Photo, error)
	DeletePhoto(c *gin.Context) (models.Photo, error)
}

func (pr *PhotoRepo) CreatePhoto(c *gin.Context) (models.Photo, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = UserID

	err := pr.db.Debug().Create(&Photo).Error

	return Photo, err
}

func (pr *PhotoRepo) GetAllPhoto(c *gin.Context) ([]models.Photo, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}
	UserID := uint(userData["id"].(float64))

	Photo.UserID = UserID
	Photo.User = &models.User{}

	var GetAllPhoto = []models.Photo{}
	err := pr.db.Model(&models.Photo{}).Find(&GetAllPhoto).Error

	return GetAllPhoto, err
}

func (pr *PhotoRepo) UpdatePhoto(c *gin.Context) (models.Photo, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := pr.db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{
		Title:    Photo.Title,
		Caption:  Photo.Caption,
		PhotoURL: Photo.PhotoURL,
	}).Error

	return Photo, err
}

func (pr *PhotoRepo) DeletePhoto(c *gin.Context) (models.Photo, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := pr.db.Exec(`
	DELETE photos 
	FROM photos 
	WHERE photos.id = ?`, photoId).Error

	return Photo, err
}
