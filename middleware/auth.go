package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

const JwtSecret = "your-secret-key"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the JWT token
		tokenString := c.GetHeader("Authorization")

		// Validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify the token signing method and return the secret key
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid token")
			}
			return []byte(JwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Token is valid, continue to the next handler
		c.Next()
	}
}
