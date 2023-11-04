package interceptor

import (
	"scaler/common/logger"
	"scaler/common/public"
	"scaler/common/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Context() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		beginTime := time.Now()
		traceId := utils.NewTraceId()
		ctx.Set(public.CtxKeyTraceId, traceId)
		logger.Info(ctx, "request",
			zap.String("method", ctx.Request.Method),
			zap.String("path", ctx.Request.URL.Path),
			zap.String("client", ctx.ClientIP()),
		)
		ctx.Next()
		logger.Info(ctx, "response",
			zap.String("method", ctx.Request.Method),
			zap.String("path", ctx.Request.URL.Path),
			zap.String("client", ctx.ClientIP()),
			zap.String("cost", time.Since(beginTime).String()),
		)
	}
}
