package user

import (
	"paper-manager/database"
)

type User struct {
	Id        int
	Username  string
	Password  string
	Salt      string
	Email     string
	State     int
	Avatar    string
	Role      string
	Signature string
}

type UserToken struct {
	Id       int
	Username string
	Email    string
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) GetUsers() (userList []User, err error) {
	if err = database.DB.Find(&userList).Error; err != nil {
		return
	}
	return
}

func (u *User) SelectByEmail(email string) (user User, err error) {
	if err = database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return
	}
	return
}
