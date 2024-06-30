package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestId() gin.HandlerFunc {
	headerXRequestID := "X-Request-Id"
	return func(ctx *gin.Context) {
		// Get id from request
		rid := ctx.GetHeader(headerXRequestID)
		if rid == "" {
			rid = uuid.New().String()
			ctx.Request.Header.Add(headerXRequestID, rid)
		}
		ctx.Set("TraceId", rid)
		// Set the id to ensure that the requestid is in the response
		ctx.Header(headerXRequestID, rid)
		ctx.Next()
	}
}
