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
	if user, err := userService.GetUserInfo(uint(userId)); err != nil {
		c.CJSON(errs.ErrDbSelect)
	} else if user == nil {
		c.CJSON(errs.ErrNoRecord, "没有用户信息")
	} else {
		c.CJSON(errs.Success, user)
	}
}

// UserDetail 获取用户详细信息
func UserDetail(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	id := ctx.Query("id")
	if len(id) == 0 {
		c.CJSON(errs.ErrParam, "用户ID")
		return
	}
	userService := service.NewUserService(ctx)
	userDetail, err := userService.GetUserList(id, 0, 1, []string{})
	if err != nil {
		c.CJSON(errs.ErrDbSelect, "用户详情")
	} else if userDetail.Count == 0 {
		c.CJSON(errs.ErrNoRecord, "用户详情")
	} else {
		c.CJSON(errs.Success, userDetail.List[0])
	}
}

// UserUnbind 解绑社交登录渠道
func UserUnbind(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	type request struct {
		Id     uint   `json:"id"`
		Source string `json:"source"`
	}
	param := &request{}
	if err := ctx.ShouldBindBodyWithJSON(param); err != nil {
		c.CJSON(errs.ErrParam, "用户id或渠道不符合要求")
		return
	}
	userService := service.NewUserService(ctx)
	if _, err := userService.UnbindUserSource(param.Id, param.Source); err != nil {
		c.CJSON(errs.ErrDbUpdate, err.Error())
		return
	}
	c.CJSON(errs.Success)

}

// UserStatus 更新用户状态
func UserStatus(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	type request struct {
		Id     uint  `json:"id"`
		Status uint8 `json:"status"`
	}
	param := &request{}
	if err := ctx.ShouldBindBodyWithJSON(param); err != nil {
		c.CJSON(errs.ErrParam, "用户id或状态值不符合要求")
		return
	}
	userService := service.NewUserService(ctx)
	if _, err := userService.UpdateUserStatus(param.Id, param.Status); err != nil {
		c.CJSON(errs.ErrDbUpdate, err.Error())
		return
	}
	c.CJSON(errs.Success)

}

// UserDelete 删除用户
func UserDelete(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	param := &struct {
		Id uint `json:"id"`
	}{}
	if err := ctx.ShouldBindBodyWithJSON(param); err != nil {
		c.CJSON(errs.ErrParam, "用户id值不符合要求")
		return
	}
	strId := ctx.Value("UserId").(string)
	if strId == strconv.Itoa(int(param.Id)) {
		c.CJSON(errs.ErrLogic, "该用户为当前账号，无法删除自己")
		return
	}
	userService := service.NewUserService(ctx)
	if _, err := userService.DeleteUser(param.Id); err != nil {
		c.CJSON(errs.ErrDbDelete, err.Error())
		return
	}
	c.CJSON(errs.Success)
}

// UserJoinGroup 加入用户组
func UserJoinGroup(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	param := &struct {
		Id    uint   `json:"id"`
		Group string `json:"group"`
	}{}
	if err := ctx.ShouldBindBodyWithJSON(param); err != nil {
		c.CJSON(errs.ErrParam, "用户id或角色值不符合要求")
		return
	}
	userService := service.NewUserService(ctx)
	if _, err := userService.JoinGroup(param.Id, param.Group); err != nil {
		c.CJSON(errs.ErrAuthGroup, err.Error())
		return
	}
	c.CJSON(errs.Success)
}

// UserExitGroup 退出用户组
func UserExitGroup(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	param := &struct {
		Id    uint   `json:"id"`
		Group string `json:"group"`
	}{}
	if err := ctx.ShouldBindBodyWithJSON(param); err != nil {
		c.CJSON(errs.ErrParam, "用户id或角色值不符合要求")
		return
	}
	userService := service.NewUserService(ctx)
	if _, err := userService.ExitGroup(param.Id, param.Group); err != nil {
		c.CJSON(errs.ErrAuthGroup, err.Error())
		return
	}
	c.CJSON(errs.Success)
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
