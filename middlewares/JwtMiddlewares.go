package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-bk/pkg/response"
	"go-bk/utils"
	"net/http"
	"strings"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			response.Error(c, http.StatusUnauthorized, 401, "未提供Token")
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, 401, "token 无效")
			return
		}

		c.Set("userId", claims.UserId)
		c.Next()
	}
}
