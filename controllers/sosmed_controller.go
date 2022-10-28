package controllers

import (
	"FP2/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SosmedController struct { // implementasi Controller
	ssa service.SosmedServiceApi
}

func NewSosmedController(ssa service.SosmedServiceApi) *SosmedController {
	return &SosmedController{ssa: ssa}
}

func (sc *SosmedController) CreateSosmedControllers(c *gin.Context) {
	res := sc.ssa.CreateSosmedService(c)
	c.JSON(http.StatusOK, res)
}

func (sc *SosmedController) GetSosmedControllers(c *gin.Context) {
	res := sc.ssa.GetAllSosmedService(c)
	c.JSON(http.StatusOK, res)
}

func (sc *SosmedController) UpdateSosmedControllers(c *gin.Context) {
	res := sc.ssa.UpdateSosmedService(c)
	c.JSON(http.StatusOK, res)
}

func (sc *SosmedController) DeleteSosmedControllers(c *gin.Context) {
	res := sc.ssa.DeleteSosmedService(c)
	c.JSON(http.StatusOK, res)
}
