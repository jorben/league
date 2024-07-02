package router

import (
	"github.com/gin-gonic/gin"
	"league/router/api"
	"net/http"
)

// SetupRouter 设置路由
func SetupRouter(s *gin.Engine) {
	// 健康检查
	s.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"ret": 0, "msg": "ok"})
	})

	s.GET("/auth/login", api.AuthLogin)
	s.GET("/auth/callback", api.AuthCallback)
	s.GET("/auth/renew", api.AuthRenew)

}
