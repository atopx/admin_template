package public

const (
	ServiceName = "scaler"
	Version     = "v0.1.0"
)

const (
	CtxKeyTraceId   = "traceId"
	CtxKeyResponse  = "response"
	CtxKeyUserId    = "userId"
	CtxKeyUserLevel = "userLevel"
)

type None struct{}

var (
	EmptyStr = ""
	EmptyObj = None{}
)
