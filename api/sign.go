package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"skr-shop-cms-api/code"
	"skr-shop-cms-api/core"
	"skr-shop-cms-api/model"
	"skr-shop-cms-api/request"
)

type UserToken struct {
	Uid string
}

func Login(context *gin.Context) (resp interface{}, err error) {
	var req request.SignIn
	if err := core.ParseRequest(context, &req); err != nil {
		fmt.Println("err", err)
	}

	data, err := model.GetStaffInfoOne("phone=?", req.Phone)
	// 正确密码验证
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))
	if err != nil {
		core.FailResp(context, code.ErrorPassWord)
	}
	return &UserToken{Uid: string(data.Id)}, nil

}
