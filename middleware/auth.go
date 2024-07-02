package middleware

import (
	"github.com/gin-gonic/gin"
	"league/common/context"
	"league/common/errs"
	"league/log"
	"league/service"
)

// Auth 登录态校验
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userId string
		c := context.CustomContext{Context: ctx}
		tokenString := ctx.GetHeader("X-Token")
		authService := service.NewAuthService(ctx)

		if len(tokenString) > 0 {
			userId, _ = authService.VerifyJwtString(tokenString)
		}
		ctx.Set("UserId", userId)

		// 校验权限
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		if isAllow := authService.IsAllow(userId, path, method); isAllow {
			log.WithField(ctx, "Path", path, "Method", method).Debugf("Check permission passed")
			ctx.Next()
		} else {
			if len(userId) > 0 {
				// 有登录 无权限
				c.CJSON(errs.ErrAuthUnauthorized)
			} else {
				// 未登录
				c.CJSON(errs.ErrAuthNoLogin)
			}
			ctx.Abort()
		}
	}
}
