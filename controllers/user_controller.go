package controllers

import (
	"FP2/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct { // implementasi Controller
	urs service.UserServiceApi
}

func NewUserController(urs service.UserServiceApi) *UserController {
	return &UserController{urs: urs}
}

func (uc *UserController) UserRegisterControllers(c *gin.Context) {
	res := uc.urs.UserRegisterService(c)
	c.JSON(http.StatusOK, res)
}

func (uc *UserController) UserLoginControllers(c *gin.Context) {
	res := uc.urs.UserLoginService(c)
	c.JSON(http.StatusOK, res)
}

func (uc *UserController) UpdateUserController(c *gin.Context) {
	res := uc.urs.UpdateUserService(c)
	c.JSON(http.StatusOK, res)
}

func (uc *UserController) DeleteUserController(c *gin.Context) {
	res := uc.urs.DeleteUserService(c)
	c.JSON(http.StatusOK, res)
}
