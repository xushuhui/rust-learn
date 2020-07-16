package api

import "github.com/gin-gonic/gin"

func IndexUser(c *gin.Context) {

}
func ShowUser(c *gin.Context) {

}
func DisableUser(c *gin.Context) {
	id := c.Param("id")
	println(id)

}
func EnableUser(c *gin.Context) {
	id := c.Param("id")
	println(id)
}
