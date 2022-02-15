package errors

type MyError struct {
	Code int
	Msg  string
	Data interface{}
}

var (
	LOGIN_UNKNOWN     = NewError(202, "用户不存在")
	LOGIN_ERROR       = NewError(203, "账号或密码错误")
	LOGIN_CHECK_ERROR = NewError(204, "输入的旧密码不匹配")
	LOGIN_DISABLE     = NewError(205, "账号暂未激活")
	VALID_ERROR       = NewError(300, "参数错误")
	FILE_TYPE_ERROR   = NewError(301, "文件类型不合法")
	FILE_UPLOAD_ERROR = NewError(302, "文件上传失败")
	ERROR             = NewError(400, "操作失败")
	UNAUTHORIZED      = NewError(401, "您还未登录")
	UNAUTHENTICATED   = NewError(402, "您权限不够")
	NOT_FOUND         = NewError(404, "资源不存在")
	INNER_ERROR       = NewError(500, "系统发生异常")
)

func (e *MyError) Error() string {
	return e.Msg
}

func NewError(code int, msg string) *MyError {
	return &MyError{
		Msg:  msg,
		Code: code,
	}
}

func GetError(e *MyError, data interface{}) *MyError {
	return &MyError{
		Msg:  e.Msg,
		Code: e.Code,
		Data: data,
	}
}
