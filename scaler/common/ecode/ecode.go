package ecode

type Ecode int

const (
	Success       Ecode = 0
	Forbidden     Ecode = 300000
	ParamError    Ecode = 400000
	Unauthorized  Ecode = 401000
	AuthExpired   Ecode = 401100
	SystemError   Ecode = 500000
	Unimplemented Ecode = 1000000
)
