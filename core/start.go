package core

import (
	"github.com/gin-gonic/gin"
)

func StartModule() {

	initRedis()
	InitValidate()
	DebugEnv()
}
func DebugEnv() {
	gin.SetMode(gin.DebugMode)

}
