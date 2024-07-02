package context

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"league/common/errs"
	"net/http"
)

type CustomContext struct {
	*gin.Context
}

type Response struct {
	Code    int         `json:"code"`    // 业务返回码
	Message string      `json:"message"` // 错误信息
	Data    interface{} `json:"data"`    // 数据
}

// CJSON 返回json结构
func (ctx *CustomContext) CJSON(code int, args ...interface{}) {
	var message string
	var data interface{}

	switch len(args) {
	case 1: // 1个传参，错误描述 或 数据
		if v, ok := args[0].(string); ok {
			message = v
			data = nil
		} else {
			data = args[0]
		}
	case 2: // 2个传参，错误描述+数据
		if v, ok := args[0].(string); ok {
			message = v
			data = args[1]
		} else if v, ok := args[1].(string); ok {
			message = v
			data = args[0]
		}
	default:
		data = nil
	}

	if len(message) > 0 {
		message = fmt.Sprintf("%s, %s", errs.GetErrorMsg(code), message)
	} else {
		message = errs.GetErrorMsg(code)
	}

	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
