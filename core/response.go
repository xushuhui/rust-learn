package core

import "github.com/gin-gonic/gin"

// JsonResponse 数据返回通用 JSON 数据结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码 ((0: 成功，1: 失败，>1: 错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据 (业务接口定义具体数据结构)
}

func ParseRequest(c *gin.Context, request interface{}) error {
	err := c.ShouldBind(request)

	if err != nil {
		c.JSON(200, JsonResponse{
			Code:    10004,
			Message: err.Error(),
		})
		return err
	}
	return nil
}
