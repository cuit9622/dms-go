package errors

type ErrorCode struct {
	Code int
	Msg  string
}

func new(code int, msg string) *ErrorCode {
	return &ErrorCode{
		Code: code,
		Msg:  msg,
	}
}

var (
	FORBIDEN              = new(403, "没有该操作权限")
	NOT_FOUND             = new(404, "请求未找到")
	INTERNAL_SERVER_ERROR = new(500, "系统异常")
)
