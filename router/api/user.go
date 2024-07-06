package api

import (
	"github.com/gin-gonic/gin"
	"league/common/context"
	"league/common/errs"
	"league/service"
	"strconv"
)

// GetUserinfo 获取当前用户信息
func GetUserinfo(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	strId := ctx.Value("UserId").(string)
	userId, err := strconv.ParseUint(strId, 10, 64) // 10进制
	if err != nil {
		c.CJSON(errs.ErrAuthNoLogin)
		return
	}
	userService := service.NewUserService(ctx)
	user := userService.GetUserInfo(uint(userId))
	if user == nil {
		c.CJSON(errs.ErrNoRecord, "没有用户信息")
		return
	}
	c.CJSON(errs.Success, user)
}
