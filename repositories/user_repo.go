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

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{
		db: db,
	}
}

type UserRepoApi interface {
	UserRegister(c *gin.Context) (models.User, error)
	UserLogin(c *gin.Context) (error, bool, string)
	UpdateUser(c *gin.Context) (models.User, models.User, error)
	DeleteUser(c *gin.Context) (models.User, error)
}

var (
	appJSON = "application/json"
)

func (ur *UserRepo) UserRegister(c *gin.Context) (models.User, error) {
	ContentType := helpers.GetContentType(c)

	var GetUser models.User

	JsonUser := GetUser
	if ContentType == appJSON {
		c.ShouldBindJSON(&JsonUser)
	} else {
		c.ShouldBind(&JsonUser)
	}
	err := ur.db.Create(&JsonUser).Error
	fmt.Println(JsonUser)
	if err != nil {
		fmt.Println(err.Error())
	}

	return JsonUser, nil
}

func (ur *UserRepo) UserLogin(c *gin.Context) (error, bool, string) {
	contentType := helpers.GetContentType(c)
	_ = contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	// Validate Email
	err := ur.db.Debug().Where("email = ?", User.Email).Take(&User).Error
	// Validate Password
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	// Validate Email & Password Jika Berhasil
	token := helpers.GenerateToken(User.ID, User.Email)

	return err, comparePass, token
}

func (ur *UserRepo) UpdateUser(c *gin.Context) (models.User, models.User, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Pengguna := models.User{}
	PenggunaDefault := models.User{}

	penggunaId, _ := strconv.Atoi(c.Param("penggunaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Pengguna)
	} else {
		c.ShouldBind(&Pengguna)
	}

	Pengguna.ID = userID           // Ambil Id dari Claims JWT
	Pengguna.ID = uint(penggunaId) // Ambil Id dari parameter

	err := ur.db.Model(&Pengguna).Where("id = ?", penggunaId).Updates(models.User{
		Email:    Pengguna.Email,
		Username: Pengguna.Username,
	}).Error

	return Pengguna, PenggunaDefault, err
}

func (ur *UserRepo) DeleteUser(c *gin.Context) (models.User, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	Pengguna := models.User{}

	penggunaId, _ := strconv.Atoi(c.Param("penggunaId"))
	userID := uint(userData["id"].(float64))

	Pengguna.ID = userID
	Pengguna.ID = uint(penggunaId)

	err := ur.db.Exec(`
	DELETE users 
	FROM users 
	WHERE users.id = ?`, penggunaId).Error

	return Pengguna, err
}
