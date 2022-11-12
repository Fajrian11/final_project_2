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

func (pc *PhotoController) CreatePhotoControllers(c *gin.Context) {
	res := pc.psa.CreatePhotoService(c)
	c.JSON(http.StatusOK, res)
}

func (pc *PhotoController) GetPhotoControllers(c *gin.Context) {
	res := pc.psa.GetAllPhotoService(c)
	c.JSON(http.StatusOK, res)
}

func (pc *PhotoController) UpdatePhotoControllers(c *gin.Context) {
	res := pc.psa.UpdatePhotoService(c)
	c.JSON(http.StatusOK, res)
}

func (pc *PhotoController) DeletePhotoControllers(c *gin.Context) {
	res := pc.psa.DeletePhotoService(c)
	c.JSON(http.StatusOK, res)
}
