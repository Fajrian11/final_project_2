package repositories

import (
	"FP2/helpers"
	"FP2/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepo {
	return CommentRepo{
		db: db,
	}
}

type CommentRepoApi interface {
	CreateComment(c *gin.Context) (models.Comment, error)
	GetAllComment(c *gin.Context) ([]models.Comment, error)
	UpdateComment(c *gin.Context) (models.Comment, error)
	DeleteComment(c *gin.Context) (models.Comment, error)
}

func (cr *CommentRepo) CreateComment(c *gin.Context) (models.Comment, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = UserID

	err := cr.db.Debug().Create(&Comment).Error

	return Comment, err
}

func (cr *CommentRepo) GetAllComment(c *gin.Context) ([]models.Comment, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comment := models.Comment{}
	UserID := uint(userData["id"].(float64))

	Comment.UserID = UserID
	Comment.User = &models.User{}

	var GetAllComment = []models.Comment{}
	err := cr.db.Preload("User").Preload("Photo").Find(&GetAllComment).Error

	return GetAllComment, err
}

func (cr *CommentRepo) UpdateComment(c *gin.Context) (models.Comment, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := cr.db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{
		Message: Comment.Message,
	}).Error

	return Comment, err
}

func (cr *CommentRepo) DeleteComment(c *gin.Context) (models.Comment, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := cr.db.Exec(`
	DELETE Comments
	FROM Comments
	WHERE Comments.id = ?`, commentId).Error

	return Comment, err
}
