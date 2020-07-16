package middleware

import (
	"github.com/gin-gonic/gin"
	"skrshop-cms-api/core"
)

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		e := c.Errors.Last()
		if e != nil {
			core.ErrorParamsResp(c, e.Err.Error())
		}

	}

	//return func(c *gin.Context) {
	//	defer func() {
	//		if err := recover(); err != nil {
	//
	//			//}else if e, ok := err.(error); ok {
	//			//	Err = OtherError(e.Error())
	//			//}else{
	//			//	Err = ServerError
	//			//}
	//			//// 记录一个错误的日志
	//			//c.JSON(Err.StatusCode,Err)
	//			fmt.Println(reflect.TypeOf(err))
	//
	//			//core.ErrorParamsResp(c,err)
	//			return
	//		}
	//	}()
	//	c.Next()
	//}
}
