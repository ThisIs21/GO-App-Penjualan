package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"app-penjualan/config"
	"app-penjualan/internal/utils"
)

func JWTAuth(cfg *config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if !strings.HasPrefix(h, "Bearer ") {
			utils.Unauthorized(c, "missing token"); c.Abort(); return
		}
		tok := strings.TrimPrefix(h, "Bearer ")
		claims := &utils.JWTClaims{}
		_, err := jwt.ParseWithClaims(tok, claims, func(token *jwt.Token) (any, error) {
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil {
			utils.Unauthorized(c, "invalid token"); c.Abort(); return
		}
		c.Set("uid", claims.UserID)
		c.Set("role", claims.Role)
		c.Set("name", claims.Name)
		c.Next()
	}

	
}

