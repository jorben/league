package middleware

import (
	"github.com/gin-gonic/gin"
	"league/service"
	"net/http"
)

// Auth 登录态校验
func Auth() gin.HandlerFunc {
	// TODO: 接入RBAC
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("x-token")
		if len(tokenString) == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"ret": -1, "msg": "x-token header is empty"})
			ctx.Abort()
		} else {
			authService := service.NewAuthService(ctx)
			if userId, err := authService.VerifyJwtString(tokenString); err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{"ret": -1, "msg": "no login"})
				ctx.Abort()
			} else {
				ctx.Set("userId", userId)
				ctx.Next()
			}
		}
		return
	}
}
