package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"league/log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		traceId := uuid.New().String()
		ctx.Set("TraceId", traceId)
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		// 记录Req
		log.WithField(
			"TraceId", traceId,
			"Method", ctx.Request.Method,
			"Path", path,
			"Ip", ctx.ClientIP(),
			"Query", query,
			"UserAgent", ctx.Request.UserAgent(),
		).Info("Request")
		ctx.Next()

		cost := time.Since(start)
		// 记录Rsp
		log.WithField(
			"TraceId", traceId,
			"Method", ctx.Request.Method,
			"Path", path,
			"Status", ctx.Writer.Status(),
			"Cost", cost,
			"Errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String(),
		).Info("Response")
	}
}
