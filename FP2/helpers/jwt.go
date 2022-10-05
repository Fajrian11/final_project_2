package helpers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

var secretKey = "token"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}
func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("Sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := string.HasPrefix(headerToken, "Bearer")
	if bearer != nil {
		return nil, errResponse
	}
	stringToken := strings.Split(headerToken, "")[1]
	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}
	return token.Claims.(jwt.MapClaims), nil
}
