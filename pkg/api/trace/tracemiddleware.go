package trace

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

const DtalkTraceId = "X-dtalk-tracd-id"

func NewTraceIdWithContext(ctx context.Context) string {
	logId, ok := ctx.Value(DtalkTraceId).(string)
	if !ok {
		logId = xid.New().String()
	}

	return logId
}

func TraceMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 或者从 header 中拿

		ctx.Set(DtalkTraceId, NewTraceIdWithContext(ctx))
		ctx.Next()
	}
}
