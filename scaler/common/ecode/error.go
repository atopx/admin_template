package ecode

type Error struct {
	Code    Ecode
	Message string
}

func New(code Ecode, message string) *Error {
	return &Error{Code: code, Message: message}
}

func NewParamError(message string) *Error {
	return &Error{Code: ParamError, Message: message}
}

func NewSystemError(message string) *Error {
	return &Error{Code: SystemError, Message: message}
}

func UnimplementedError() error {
	return &Error{Code: Unimplemented, Message: "Unimplemented"}
}

func (e *Error) Error() string {
	return e.Message
}
