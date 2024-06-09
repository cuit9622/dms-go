package errors

type ErrorCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func new(code int, msg string) *ErrorCode {
	return &ErrorCode{
		Code: code,
		Msg:  msg,
	}
}

var (
	BAD_REQUEST           = new(400, "请求参数不正确")
	FORBIDEN              = new(403, "没有该操作权限")
	NOT_FOUND             = new(404, "请求未找到")
	INTERNAL_SERVER_ERROR = new(500, "系统异常")
)
