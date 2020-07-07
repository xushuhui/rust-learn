package core

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
)

func InitValidate() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		// 自定义验证方法
		_ = v.RegisterValidation("isMobile", isMobile)

	}
}

func isMobile(fl validator.FieldLevel) bool {
	mobile := strconv.Itoa(int(fl.Field().Uint()))
	re := `^1[3456789]\d{9}$`
	r := regexp.MustCompile(re)
	return r.MatchString(mobile)
}
