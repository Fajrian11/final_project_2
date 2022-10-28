package controllers

import (
	"FP2/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController struct { // implementasi Controller
	psa service.CommentServiceApi
}

func NewCommentController(psa service.CommentServiceApi) *CommentController {
	return &CommentController{psa: psa}
}

func (uc *CommentController) CreateCommentControllers(c *gin.Context) {
	res := uc.psa.CreateCommentService(c)
	c.JSON(http.StatusOK, res)
}

func (uc *CommentController) GetCommentControllers(c *gin.Context) {
	res := uc.psa.GetAllCommentService(c)
	c.JSON(http.StatusOK, res)
}

func (uc *CommentController) UpdateCommentControllers(c *gin.Context) {
	res := uc.psa.UpdateCommentService(c)
	c.JSON(http.StatusOK, res)
}

func (uc *CommentController) DeleteCommentControllers(c *gin.Context) {
	res := uc.psa.DeleteCommentService(c)
	c.JSON(http.StatusOK, res)
}
