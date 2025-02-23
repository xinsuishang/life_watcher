package errno

import "fmt"

const (
	SuccessCode    = 0
	ParamErrCode   = 400
	AuthErrCode    = 401
	ServiceErrCode = 500
)

var (
	DataExistCode = NewErrorCode(ParamErrCode, 001)
)

const (
	SuccessMsg    = "Success"
	ParamErrMsg   = "params error"
	AuthErrMsg    = "authorization failed"
	ServiceErrMsg = "service error"
)

func NewErrorCode(parentCode, code int32) int32 {
	return parentCode*100 + code
}

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

var (
	Success                = NewErrNo(SuccessCode, SuccessMsg)
	ParamErr               = NewErrNo(ParamErrCode, ParamErrMsg)
	AuthorizationFailedErr = NewErrNo(AuthErrCode, AuthErrMsg)
)
