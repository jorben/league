package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupRouter 设置路由
func SetupRouter(s *gin.Engine) {
	// 健康检查
	s.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"ret": 0, "msg": "ok"})
	})
}
