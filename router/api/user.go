package api

import (
	"github.com/gin-gonic/gin"
	"league/common"
	"league/common/context"
	"league/common/errs"
	"league/service"
	"strconv"
)

// UserCurrent 获取当前用户信息
func UserCurrent(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	strId := ctx.Value("UserId").(string)
	userId, err := strconv.ParseUint(strId, 10, 64) // 10进制
	if err != nil || userId == 0 {
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

// UserList 获取用户列表
func UserList(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	offset := 0
	limit := common.DEFAULT_PAGESIZE
	order := []string{"ID desc"}
	searchKey := ctx.Query("search")

	if v, err := strconv.ParseInt(ctx.Query("size"), 10, 64); err == nil {
		limit = int(v)
		if limit <= 0 {
			limit = common.DEFAULT_PAGESIZE
		}
	}
	if v, err := strconv.ParseInt(ctx.Query("page"), 10, 64); err == nil {
		offset = (int(v) - 1) * limit
		if offset < 0 {
			offset = 0
		}
	}

	userService := service.NewUserService(ctx)
	userList, err := userService.GetUserList(searchKey, offset, limit, order)
	if err != nil {
		c.CJSON(errs.ErrDbSelect, "用户列表")
		return
	}
	c.CJSON(errs.Success, userList)
}
