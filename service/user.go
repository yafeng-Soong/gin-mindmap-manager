package service

import (
	"github.com/yafeng-Soong/gin-mindmap-manager/global"
	"github.com/yafeng-Soong/gin-mindmap-manager/mapper"
	"github.com/yafeng-Soong/gin-mindmap-manager/model/errors"
	"github.com/yafeng-Soong/gin-mindmap-manager/model/user"
	"github.com/yafeng-Soong/gin-mindmap-manager/model/user/request"
	"github.com/yafeng-Soong/gin-mindmap-manager/utils"
	"go.uber.org/zap"
)

type UserService struct{}

var userMapper mapper.UserMapper

func (s *UserService) Login(login *request.RegisterAndLogin) (*user.User, error) {
	var u user.User
	var err error
	u, err = userMapper.SelectByEmail(login.Email)
	if err != nil {
		global.LOG.Info("登陆错误", zap.Error(err))
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
