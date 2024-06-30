package middleware

import (
	"github.com/gin-gonic/gin"
	"league/log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		log.WithField(
			ctx,
			"Method", ctx.Request.Method,
			"Path", ctx.Request.URL.Path,
			"Query", ctx.Request.URL.RawQuery,
			"Ip", ctx.ClientIP(),
			//"UserAgent", ctx.Request.UserAgent(),
		).Info("Request")

		ctx.Next()

		cost := time.Since(start)
		// 记录Rsp
		log.WithField(
			ctx,
			"Method", ctx.Request.Method,
			"Path", ctx.Request.URL.Path,
			"Status", ctx.Writer.Status(),
			"Cost", cost,
			"Errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String(),
		).Info("Response")
	}
}
