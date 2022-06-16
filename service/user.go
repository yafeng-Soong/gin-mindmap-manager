package service

import (
	"log"

	"github.com/yafeng-Soong/gin-mindmap-manager/model/errors"
	"github.com/yafeng-Soong/gin-mindmap-manager/model/user"
	"github.com/yafeng-Soong/gin-mindmap-manager/model/user/request"
	"github.com/yafeng-Soong/gin-mindmap-manager/utils"
)

type UserService struct{}

var userModel user.User

func (s *UserService) Login(login *request.RegisterAndLogin) (*user.User, error) {
	var u user.User
	var err error
	u, err = userModel.SelectByEmail(login.Email)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.LOGIN_UNKNOWN
	}
	if u.State == 0 {
		return nil, errors.LOGIN_DISABLE
	}
	passwd := utils.MD5V([]byte(login.Password), []byte(u.Salt), utils.TIMES)
	if passwd != u.Password {
		return nil, errors.LOGIN_ERROR
	}
	return &u, nil
}
