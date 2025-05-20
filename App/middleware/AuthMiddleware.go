package middleware

import (
	"Golunwen2/App/model"
	"Golunwen2/App/tool"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(401, gin.H{"error": "未授权的请求"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:] // 去掉 "Bearer " 前缀
		token, claims, err := tool.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "权限不足"})
			c.Abort()
			return
		}

		//验证通过后获取claims中的用户信息
		userId := claims.UserId
		var user model.User
		model.DB.First(&user, userId)

		//验证用户是否存在
		if user.Id == 0 {
			c.JSON(401, gin.H{"error": "用户不存在"})
			c.Abort()
			return
		}
		//将用户信息存入上下文
		c.Set("user", user)
		c.Next()
	}
}
