package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	errCode "skr-shop-cms-api/code"
)

// JsonResponse 数据返回通用 JSON 数据结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码 ((0: 成功，1: 失败，>1: 错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据 (业务接口定义具体数据结构)
}

func ParseRequest(c *gin.Context, request interface{}) error {
	err := c.ShouldBind(request)

	var errStr string
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errStr = Translate(err.(validator.ValidationErrors))
		case *json.UnmarshalTypeError:
			unmarshalTypeError := err.(*json.UnmarshalTypeError)
			errStr = fmt.Errorf("%s 类型错误，期望类型 %s", unmarshalTypeError.Field, unmarshalTypeError.Type.String()).Error()
		default:
			errStr = errors.New("unknown error").Error()
		}

		ErrorParamsResp(c, errStr)
		return err
	}
	return nil
}
func FailResp(c *gin.Context, code int) {
	c.AbortWithStatusJSON(200, JsonResponse{
		Code:    code,
		Message: errCode.GetMsg(code),
	})
	return
}
func ErrorParamsResp(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(200, JsonResponse{
		Code:    errCode.InvalidParams,
		Message: msg,
	})
	return
}
func SuccessResp(c *gin.Context) {
	c.JSON(200, JsonResponse{
		Code:    0,
		Message: errCode.GetMsg(0),
	})
}
func SetData(c *gin.Context, data interface{}) {
	c.JSON(200, JsonResponse{
		Code:    0,
		Message: errCode.GetMsg(0),
		Data:    data,
	})
}
