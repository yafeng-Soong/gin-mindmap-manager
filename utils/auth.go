package utils

import (
	"paper-manager/model/user/response"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type serverError struct {
	msg string
}

func (se *serverError) Error() string {
	return se.msg
}

func GetCurrentUser(c *gin.Context) (*response.UserInfo, error) {
	auth := c.Request.Header.Get("Authorization")
	claims, err := ParseToken(auth)
	userInfo := &response.UserInfo{Id: claims.User.Id, Username: claims.User.Username, Email: claims.User.Email}
	return userInfo, err
}

// 使用Session获取当前登录用户
func GetCurrentUser1(c *gin.Context) (userInfo *response.UserInfo, err error) {
	session := sessions.Default(c)
	info := session.Get("currentUser").(response.UserInfo)
	userInfo = &info
	if userInfo == nil {
		return nil, &serverError{msg: "未登录"}
	} else {
		return userInfo, nil
	}
}
