package interceptor

import (
	"app/common/ecode"
	"app/common/public"
	"app/internal/model"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth(level int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err *ecode.Error
		defer func() {
			if err != nil {
				traceId := ctx.GetString(public.CtxKeyTraceId)
				ctx.AbortWithStatusJSON(http.StatusOK, public.NewErrorResponse(traceId, err))
			} else {
				ctx.Next()
			}
		}()

		token := ctx.Request.Header.Get("Authorization")
		if token == public.EmptyStr || !strings.HasPrefix(token, "Bearer ") {
			err = ecode.New(ecode.Unauthorized, "用户未登录")
			return
		}

		tokenStr := token[strings.Index(token, " ")+1:]

		userToken := new(model.UserToken)

		if err := userToken.First("access_token=?", tokenStr); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = ecode.New(ecode.Unauthorized, "用户未登录")
				return
			}
			err = ecode.NewSystemError(err.Error())
			return
		}

		if userToken.ExpireTime < time.Now().Local().Unix() {
			err = ecode.New(ecode.AuthExpired, "用户认证过期")
			return
		}

		user := new(model.User)
		if err := user.First("id=?", userToken.Id); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = ecode.New(ecode.Unauthorized, "用户不存在")
				return
			}
			err = ecode.NewSystemError(err.Error())
			return
		}

		if user.Status == model.UserStatusDisable {
			err = ecode.New(ecode.Unauthorized, "用户已被禁用")
			return
		}

		if user.Level < level {
			err = ecode.New(ecode.Forbidden, "权限不足, 禁止访问")
			return
		}

		ctx.Set(public.CtxKeyUserId, user.Id)
		ctx.Set(public.CtxKeyUserLevel, user.Level)
	}
}
