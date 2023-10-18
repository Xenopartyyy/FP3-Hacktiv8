package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretkey = "rahasia"
var err error

func GenerateToken(ID uint, Email string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    ID,
		"email": Email,
	})

	signedToken, err := token.SignedString([]byte(secretkey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, err
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, err := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(secretkey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && token.Valid {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}
