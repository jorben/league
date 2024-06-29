package router

import (
	"github.com/gin-gonic/gin"
	"league/middleware"
	"league/router/api/auth"
	"net/http"
)

// SetupRouter 设置路由
func SetupRouter(s *gin.Engine) {
	// 健康检查
	s.GET("/health", middleware.Auth(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"ret": 0, "msg": "ok"})
	})

	s.GET("/auth/github_login", auth.LoginGithub)

	s.GET("/auth/github_callback", auth.CallbackGithub)

}
