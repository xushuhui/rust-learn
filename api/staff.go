package api

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"skrshop-cms-api/core"
	"skrshop-cms-api/model"
	"skrshop-cms-api/request"
)

func StoreStaff(c *gin.Context) {
	var req request.Staff

	if err := core.ParseRequest(c, &req); err != nil {
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Error(err)
		return
	}
	staff := model.StaffInfo{Name: req.Name, Email: req.Email, Password: string(password), Phone: req.Phone}
	err = staff.Create()
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return

}
func IndexStaff(c *gin.Context) {
	data := model.GetStaffInfoPage(1, 10)
	core.SetData(c, data)

	return
}
func ShowStaff(c *gin.Context) {
	id := c.Param("id")
	data, err := model.GetStaffInfoById(id)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)

	return
}
func DestroyStaff(c *gin.Context) {
	return
}
func UpdateStaff(c *gin.Context) {
	return
}
