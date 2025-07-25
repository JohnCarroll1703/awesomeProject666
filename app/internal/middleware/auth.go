package middleware

import (
	"awesomeProject666/app/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTAuth(jwtManager *auth.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwtManager.Verify(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("userID", claims.UserID)

		c.Next()
	}
}
