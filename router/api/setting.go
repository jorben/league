package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"league/common"
	"league/common/context"
	"league/common/errs"
	"league/model"
	"league/service"
	"strconv"
)

// SettingApiList 获取Api列表
func SettingApiList(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	offset := 0
	limit := common.DEFAULT_PAGESIZE
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
	settingService := service.NewSettingService(ctx)
	if list, err := settingService.GetApiList(offset, limit); err != nil {
		c.CJSON(errs.ErrDbSelect)
	} else {
		c.CJSON(errs.Success, list)
	}
}

// SettingUpdateApi 创建/更新Api
func SettingUpdateApi(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	param := &model.Api{}
	if err := ctx.ShouldBindBodyWithJSON(param); err != nil {
		c.CJSON(errs.ErrParam, "Api信息不符合要求")
		return
	}
	settingService := service.NewSettingService(ctx)
	_, err := settingService.SaveApi(param)
	if err != nil {
		c.CJSON(errs.ErrDbUpdate)
		return
	}
	c.CJSON(errs.Success)
}

// SettingDeleteApi 删除API
func SettingDeleteApi(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	param := &struct {
		ID uint `json:"ID"`
	}{}
	if err := ctx.ShouldBindBodyWithJSON(param); err != nil || param.ID == 0 {
		c.CJSON(errs.ErrParam, "Api ID不符合要求")
		return
	}
	settingService := service.NewSettingService(ctx)
	err := settingService.DeleteApi(param.ID)
	if err != nil {
		c.CJSON(errs.ErrDbDelete, "删除接口失败")
		return
	}
	c.CJSON(errs.Success)
}

// SettingMenuList 获取全部菜单
func SettingMenuList(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	menuService := service.NewMenuService(ctx)
	menus, err := menuService.GetAllMenus()
	if err != nil {
		c.CJSON(errs.ErrDbSelect, "菜单查询失败")
		return
	}
	c.CJSON(errs.Success, menus)
}

// SettingUpdateMenu 创建/更新菜单项
func SettingUpdateMenu(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	param := &model.Menu{}
	if err := ctx.ShouldBindBodyWithJSON(param); err != nil {
		c.CJSON(errs.ErrParam, "菜单项信息不符合要求")
		return
	}
	if len(param.Key) == 0 || len(param.Type) == 0 || len(param.Label) == 0 {
		c.CJSON(errs.ErrParam, "缺少必要的菜单信息")
		return
	}
	menuService := service.NewMenuService(ctx)
	_, err := menuService.SaveMenu(param)
	if err != nil {
		if errors.Is(err, errs.ErrorRecordNotFound) {
			c.CJSON(errs.ErrLogic, "菜单父节点不存在，请检查")
			return
		}
		c.CJSON(errs.ErrDbUpdate)
		return
	}
	c.CJSON(errs.Success)
}

// SettingDeleteMenu 删除菜单项
func SettingDeleteMenu(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	param := &struct {
		ID uint `json:"ID"`
	}{}
	if err := ctx.ShouldBindBodyWithJSON(param); err != nil || param.ID == 0 {
		c.CJSON(errs.ErrParam, "菜单项ID不符合要求")
		return
	}
	menuService := service.NewMenuService(ctx)
	err := menuService.DeleteMenu(param.ID)
	if err != nil {
		if errors.Is(err, errs.ErrorHasChildren) {
			c.CJSON(errs.ErrLogic, "该菜单含有子菜单，无法直接删除")
			return
		}
		c.CJSON(errs.ErrDbDelete, "删除菜单项失败")
		return
	}
	c.CJSON(errs.Success)
}
