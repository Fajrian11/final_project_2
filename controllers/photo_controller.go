package controllers

import (
	"FP2/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhotoController struct { // implementasi Controller
	psa service.PhotoServiceApi
}

func NewPhotoController(psa service.PhotoServiceApi) *PhotoController {
	return &PhotoController{psa: psa}
}

func (uc *PhotoController) CreatePhotoControllers(c *gin.Context) {
	res := uc.psa.CreatePhotoService(c)
	c.JSON(http.StatusOK, res)
}

func (uc *PhotoController) GetPhotoControllers(c *gin.Context) {
	res := uc.psa.GetAllPhotoService(c)
	c.JSON(http.StatusOK, res)
}

func (uc *PhotoController) UpdatePhotoControllers(c *gin.Context) {
	res := uc.psa.UpdatePhotoService(c)
	c.JSON(http.StatusOK, res)
}

func (uc *PhotoController) DeletePhotoControllers(c *gin.Context) {
	res := uc.psa.DeletePhotoService(c)
	c.JSON(http.StatusOK, res)
}
