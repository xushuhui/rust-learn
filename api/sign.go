package api

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"skrshop-cms-api/code"
	"skrshop-cms-api/core"
	"skrshop-cms-api/model"
	"skrshop-cms-api/request"
)

type UserToken struct {
	Uid string
}

func Login(c *gin.Context) {
	var req request.SignIn
	if err := core.ParseRequest(c, &req); err != nil {
		return
	}

	data, err := model.GetStaffInfoOne("phone=?", req.Phone)
	// 正确密码验证
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))
	if err != nil {
		core.FailResp(c, code.ErrorPassWord)
	}

}
