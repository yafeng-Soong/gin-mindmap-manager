package mapper

import (
	"github.com/yafeng-Soong/gin-mindmap-manager/global"
	"github.com/yafeng-Soong/gin-mindmap-manager/model/user"
)

type UserMapper struct{}

func (u *UserMapper) GetUsers() (userList []user.User, err error) {
	if err = global.DB.Find(&userList).Error; err != nil {
		return
	}
	return
}

func (u *UserMapper) SelectByEmail(email string) (user user.User, err error) {
	if err = global.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return
	}
	return
}
