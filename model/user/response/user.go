package response

import (
	"paper-manager/model/user"
)

type UserInfo struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Signature string `json:"signature"`
	Role      string `json:"role"`
}

type LoginResponseVo struct {
	User  UserInfo `json:"userInfo"`
	Token string   `json:"token"`
}

func NewUserInfo(user user.User) *UserInfo {
	return &UserInfo{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Signature: user.Signature,
		Role:      user.Role,
	}
}

func NewLoginResponse(user user.User, token string) *LoginResponseVo {

	userInfo := &UserInfo{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Signature: user.Signature,
		Role:      user.Role,
	}
	return &LoginResponseVo{
		User:  *userInfo,
		Token: token,
	}
}
