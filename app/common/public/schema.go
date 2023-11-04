package public

import "app/common/ecode"

type Response struct {
	Code    ecode.Ecode `json:"code"`
	Message string      `json:"message"`
	TraceId string      `json:"traceId"`
	Data    any         `json:"data"`
}

func NewResponse(traceId string, data any) (resp *Response) {
	resp = new(Response)
	resp.Code = ecode.Success
	resp.TraceId = traceId
	resp.Data = data
	return resp
}

func NewErrorResponse(traceId string, err *ecode.Error) (resp *Response) {
	resp = new(Response)
	resp.Code = err.Code
	resp.Message = err.Message
	resp.TraceId = traceId
	return resp
}

func NewSystemErrorResponse(traceId, message string) (resp *Response) {
	return NewErrorResponse(traceId, ecode.NewSystemError(message))
}

type PageInfo struct {
	Index    int   `json:"index"`
	Size     int   `json:"size"`
	Count    int64 `json:"count"`
	Disabled bool  `json:"disabled"`
}

type TimeRange struct {
	Left  int64 `json:"left"`
	Right int64 `json:"right"`
}
