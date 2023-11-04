package controller

import (
	"net/http"
	"scaler/common/ecode"
	"scaler/common/logger"
	"scaler/common/public"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Controller 控制器
type Controller struct {
	traceId string
	context *gin.Context
	Params  any
	err     error
}

func New(ctx *gin.Context, params any) *Controller {
	err := ctx.ShouldBind(params)
	if err != nil {
		err = ecode.NewParamError(err.Error())
	}
	return &Controller{
		traceId: ctx.GetString(public.CtxKeyTraceId),
		context: ctx,
		Params:  params,
		err:     err,
	}
}

func (ctl *Controller) UserId() int {
	return ctl.context.GetInt(public.CtxKeyUserId)
}

func (ctl *Controller) Deal() (any, error) {
	return nil, ecode.UnimplementedError()
}

func (ctl *Controller) Error() error {
	return ctl.err
}

func (ctl *Controller) Context() *gin.Context { return ctl.context }

func (ctl *Controller) About(data any, err error) {
	var resp *public.Response
	if err != nil {
		if cerr, ok := err.(*ecode.Error); ok {
			resp = public.NewErrorResponse(ctl.traceId, cerr)
		} else {
			resp = public.NewSystemErrorResponse(ctl.traceId, err.Error())
		}
	} else {
		resp = public.NewResponse(ctl.traceId, data)
	}
	ctl.context.JSON(http.StatusOK, resp)
}

type Interface interface {
	Deal() (any, error)
	Context() *gin.Context
	Error() error
	About(any, error)
}

func Handle(ctl Interface) {
	if err := ctl.Error(); err != nil {
		logger.Warn(ctl.Context(), "bind param error", zap.Error(err))
		ctl.About(nil, err)
		return
	}
	ctl.About(ctl.Deal())
}
