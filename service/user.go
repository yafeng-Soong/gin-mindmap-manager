package service

import (
	"log"
	"paper-manager/model/errors"
	"paper-manager/model/user"
	"paper-manager/model/user/request"
	"paper-manager/utils"
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
