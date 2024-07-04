package router

import (
	"github.com/gin-gonic/gin"
	"league/common/context"
	"league/common/errs"
	"league/middleware"
	"league/router/api"
)

// SetupRouter 设置路由
func SetupRouter(s *gin.Engine) {
	// 健康检查
	s.GET("/health", func(ctx *gin.Context) {
		c := context.CustomContext{Context: ctx}
		c.CJSON(errs.Success)
	})

	backend := s.Group("/api")
	backend.Use(middleware.RequestId(), middleware.Logger(), middleware.Auth(), gin.Recovery())
	backend.GET("/auth/login", api.AuthLogin)
	backend.GET("/auth/callback", api.AuthCallback)
	backend.GET("/auth/renew", api.AuthRenew)
	backend.GET("/auth/logout", api.AuthLogout)

	s.Static("/static", "./web/build/static")
	s.NoRoute(func(ctx *gin.Context) {
		ctx.File("./web/build/index.html")
	})

}
