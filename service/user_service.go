package service

import (
	"FP2/repositories"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	rr repositories.UserRepoApi
}

func NewUserService(rr repositories.UserRepoApi) *UserService { //provie service
	return &UserService{rr: rr}
}

type UserServiceApi interface {
	UserRegisterService(c *gin.Context) gin.H
	UserLoginService(c *gin.Context) gin.H
	UpdateUserService(c *gin.Context) gin.H
	DeleteUserService(c *gin.Context) gin.H
}

func (s UserService) UserRegisterService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	user, err := s.rr.UserRegister(c)
	if err != nil {
		result = gin.H{
			"result": "Failed Create User",
		}
	} else {
		result = gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"age":      user.Age,
		}
	}
	return result
}

func (s UserService) UserLoginService(c *gin.Context) gin.H {
	var result gin.H

	err, comparePass, token := s.rr.UserLogin(c)

	// Validate Email
	if err != nil {
		result = gin.H{
			"error":   "Unauthorized",
			"message": "invalid email / password",
		}
	}
	// Validate Password
	if !comparePass {
		result = gin.H{
			"error":   "Unauthorized",
			"message": "invalid email / password",
		}
	}
	// Validate Email & Password Jika Berhasil
	if err == nil && comparePass {
		result = gin.H{
			"token": token,
		}
	}

	return result
}

func (us UserService) UpdateUserService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	Pengguna, PenggunaDefault, err := us.rr.UpdateUser(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"Success":    "Data Has been Updated",
			"id":         Pengguna.ID,
			"email":      PenggunaDefault.Email,
			"username":   Pengguna.Username,
			"age":        Pengguna.Age,
			"updated_at": Pengguna.UpdatedAt,
		}
	}
	return result
}

func (us UserService) DeleteUserService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	_, err := us.rr.DeleteUser(c)
	if err != nil {
		result = gin.H{
			"result":  "Gagal Menghapus Data",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"Success": "Your account has been successfully deleted",
		}
	}
	return result
}
