package auth

import (
	"net/http"
	"pro/internal/global"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token from cookie
		tokenString, _ := c.Cookie("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return global.JWT_KEY, nil 
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token BBB"})
			c.Abort()
			return
		}
  
		c.Next()
	}
}