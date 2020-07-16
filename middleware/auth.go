package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"skrshop-cms-api/core"
	"time"
)

func AuthMiddleware() {

}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)

		m := claims[jwt.IdentityKey]
		token := m.(map[string]interface{})

		c.Set("uid", token["uid"])
		c.Next()
	}
}
func jwtUnAuthFunc(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
func loginResponse(c *gin.Context, code int, token string, expire time.Time) {
	m := make(map[string]interface{})
	m["token"] = token
	m["expire"] = expire
	core.SetData(c, m)
}
