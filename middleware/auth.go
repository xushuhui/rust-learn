package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"skr-shop-cms-api/api"
	"skr-shop-cms-api/core"
	"time"
)

func AuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "Realmname",
		Key:           []byte("Secretkey"),
		Timeout:       time.Hour * 12,
		MaxRefresh:    time.Hour * 24,
		Authenticator: api.Login,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			return jwt.MapClaims{
				jwt.IdentityKey: data,
			}

		},
		Unauthorized:  jwtUnAuthFunc,
		LoginResponse: loginResponse,

		// 其他默认
	})
	if err != nil {
		return nil
	}
	return authMiddleware

}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)

		m := claims[jwt.IdentityKey]
		token := m.(map[string]interface{})

		c.Set("uid", token["uid"])
		//_ = utils.MapToStruct(claims, &token)
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
