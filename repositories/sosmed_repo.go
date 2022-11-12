package repositories

import (
	"FP2/helpers"
	"FP2/models"
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type SosmedRepo struct {
	db *gorm.DB
}

func NewSosmedRepo(db *gorm.DB) SosmedRepo {
	return SosmedRepo{
		db: db,
	}
}

type SosmedRepoApi interface {
	CreateSosmed(c *gin.Context) (models.SocialMedia, error)
	GetAllSosmed(c *gin.Context) ([]models.SocialMedia, error)
	UpdateSosmed(c *gin.Context) (models.SocialMedia, error)
	DeleteSosmed(c *gin.Context) (models.SocialMedia, error)
}

func (sr *SosmedRepo) CreateSosmed(c *gin.Context) (models.SocialMedia, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Sosmed := models.SocialMedia{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Sosmed)
	} else {
		c.ShouldBind(&Sosmed)
	}

	Sosmed.UserID = UserID

	err := sr.db.Debug().Create(&Sosmed).Error

	return Sosmed, err
}

func (sr *SosmedRepo) GetAllSosmed(c *gin.Context) ([]models.SocialMedia, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	Sosmed := models.SocialMedia{}
	UserID := uint(userData["id"].(float64))

	Sosmed.UserID = UserID
	Sosmed.User = &models.User{}

	var GetAllSosmed = []models.SocialMedia{}
	// err := sr.db.Model(&models.SocialMedia{}).Find(&GetAllSosmed).Error
	err := sr.db.Preload("User").Find(&GetAllSosmed).Error
	if err != nil {
		fmt.Println(err.Error())
	}

	return GetAllSosmed, err
}

func (sr *SosmedRepo) UpdateSosmed(c *gin.Context) (models.SocialMedia, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Sosmed := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Sosmed)
	} else {
		c.ShouldBind(&Sosmed)
	}

	Sosmed.UserID = userID
	Sosmed.ID = uint(socialMediaId)

	err := sr.db.Model(&Sosmed).Where("id = ?", socialMediaId).Updates(models.SocialMedia{
		Name:           Sosmed.Name,
		SocialMediaURL: Sosmed.SocialMediaURL,
	}).Error

	return Sosmed, err
}

func (pr *SosmedRepo) DeleteSosmed(c *gin.Context) (models.SocialMedia, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	Sosmed := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	Sosmed.UserID = userID
	Sosmed.ID = uint(socialMediaId)

	err := pr.db.Exec(`
	DELETE social_media
	FROM social_media
	WHERE social_media.id = ?`, socialMediaId).Error

	return Sosmed, err
}
