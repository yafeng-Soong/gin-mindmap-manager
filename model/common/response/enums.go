package response

type ResponseEnums struct {
	Code int
	Msg  string
}

func newEnums(code int, msg string) *ResponseEnums {
	return &ResponseEnums{
		Code: code,
		Msg:  msg,
	}
}

func SUCCESS() *ResponseEnums {
	return newEnums(200, "操作成功")
}

func LOGIN_SUCCESS() *ResponseEnums {
	return newEnums(201, "登录成功")
}

func LOGIN_UNKNOWN() *ResponseEnums {
	return newEnums(202, "用户不存在")
}

func LOGIN_ERROR() *ResponseEnums {
	return newEnums(203, "账号或密码错误")
}

func LOGIN_CHECK_ERROR() *ResponseEnums {
	return newEnums(204, "输入的旧密码不匹配")
}

func LOGIN_DISABLE() *ResponseEnums {
	return newEnums(205, "账号暂未激活")
}

func VALID_ERROR() *ResponseEnums {
	return newEnums(300, "参数错误")
}

func FILE_TYPE_ERROR() *ResponseEnums {
	return newEnums(301, "文件类型不合法")
}

func FILE_UPLOAD_ERROR() *ResponseEnums {
	return newEnums(302, "文件上传失败")
}

func ERROR() *ResponseEnums {
	return newEnums(400, "操作失败")
}

func UNAUTHORIZED() *ResponseEnums {
	return newEnums(401, "您还未登录")
}

func UNAUTHENTICATED() *ResponseEnums {
	return newEnums(402, "您权限不够")
}

func NOT_FOUND() *ResponseEnums {
	return newEnums(404, "资源不存在")
}

func INNER_ERROR() *ResponseEnums {
	return newEnums(500, "系统发生异常")
}
