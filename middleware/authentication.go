package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequiredAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}
