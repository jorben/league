package router

import (
	"github.com/gin-gonic/gin"
	"league/common/context"
	"league/common/errs"
	"league/router/api"
)

// SetupRouter 设置路由
func SetupRouter(s *gin.Engine) {
	// 健康检查
	s.GET("/health", func(ctx *gin.Context) {
		c := context.CustomContext{Context: ctx}
		c.CJSON(errs.Success)
	})

	s.GET("/auth/login", api.AuthLogin)
	s.GET("/auth/callback", api.AuthCallback)
	s.GET("/auth/renew", api.AuthRenew)
	s.GET("/auth/logout", api.AuthLogout)

}
