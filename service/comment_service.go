package service

import (
	"FP2/repositories"

	"github.com/gin-gonic/gin"
)

type CommentService struct {
	rr repositories.CommentRepoApi
}

func NewCommentService(rr repositories.CommentRepoApi) *CommentService { //provie service
	return &CommentService{rr: rr}
}

type CommentServiceApi interface {
	CreateCommentService(c *gin.Context) gin.H
	GetAllCommentService(c *gin.Context) gin.H
	UpdateCommentService(c *gin.Context) gin.H
	DeleteCommentService(c *gin.Context) gin.H
}

func (cs CommentService) CreateCommentService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	Comment, _ := cs.rr.CreateComment(c)
	if Comment.Message == "" {
		result = gin.H{
			"error": "Your message is required",
		}
	} else if Comment.PhotoID == 0 {
		result = gin.H{
			"error": "Your photo_id is required",
		}
	} else {
		result = gin.H{
			"Success":    "Data Has been created",
			"id":         Comment.ID,
			"message":    Comment.Message,
			"photo_id":   Comment.PhotoID,
			"user_id":    Comment.UserID,
			"created_at": Comment.CreatedAt,
		}
	}
	return result
}

func (cs CommentService) GetAllCommentService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	GetAllComment, err := cs.rr.GetAllComment(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"result": GetAllComment,
			"count":  len(GetAllComment),
		}
	}
	return result
}

func (cs CommentService) UpdateCommentService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	Comment, err := cs.rr.UpdateComment(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"Success":    "Data Has been Updated",
			"id":         Comment.ID,
			"message":    Comment.Message,
			"photo_id":   Comment.PhotoID,
			"user_id":    Comment.UserID,
			"created_at": Comment.UpdatedAt,
		}
	}
	return result
}

func (cs CommentService) DeleteCommentService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	_, err := cs.rr.DeleteComment(c)
	if err != nil {
		result = gin.H{
			"result":  "Gagal Menghapus Data",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"Success": "Your Comment has been successfully deleted",
		}
	}
	return result
}
