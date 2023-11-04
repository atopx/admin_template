package interceptor

import (
	"app/common/logger"
	"app/common/public"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				body, _ := httputil.DumpRequest(ctx.Request, false)
				if brokenPipe {
					logger.Error(ctx, "recovery from broken pipeline",
						zap.ByteString("body", body),
						zap.Any("error", err),
					)
					if err != nil {
						panic(err)
					}
					ctx.Abort()
					return
				}
				stack := zap.Stack("stack")
				logger.Error(ctx, "recovery from panic",
					zap.Any("error", err),
					zap.ByteString("body", body),
					stack,
				)
				traceId := ctx.GetString(public.CtxKeyTraceId)
				fmt.Println("========================[PANIC]========================")
				fmt.Println("time:", time.Now().Local().Format(time.DateTime))
				fmt.Println("traceId: ", traceId)
				fmt.Println(stack.String)
				resp := public.NewSystemErrorResponse(traceId, "InternalServerError")
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
			}
		}()
		ctx.Next()
	}
}
