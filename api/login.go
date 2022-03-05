package api

import (
	"log"
	"paper-manager/model/common/response"
	"paper-manager/model/errors"
	"paper-manager/model/user"
	"paper-manager/model/user/request"
	user_response "paper-manager/model/user/response"
	"paper-manager/service"
	"paper-manager/utils"

	"github.com/gin-gonic/gin"
)

type LoginApi struct {
}

var userModel user.User
var userService service.UserService

// 测试接口获得所有用户列表
func (l *LoginApi) GetUsers(c *gin.Context) {
	users, err := userModel.GetUsers()
	if err != nil {
		c.Error(err)
		return
	}
	// session := sessions.Default(c)
	// log.Println(session.Get("currentUser"))
	response.OkWithData(users, c)
}

// 测试接口获得指定email用户
func (l *LoginApi) GetUserByEmial(c *gin.Context) {
	email := c.Query("email")
	user, err := userModel.SelectByEmail(email)
	if err != nil {
		c.Error(err)
		return
	}
	response.OkWithData(user, c)
}

func (l *LoginApi) Login(c *gin.Context) {
	var loginVo request.RegisterAndLogin
	if e := c.ShouldBindJSON(&loginVo); e != nil {
		log.Println(e.Error())
		c.Error(errors.VALID_ERROR)
		return
	}
	u, err := userService.Login(&loginVo)
	if err != nil {
		c.Error(err)
		// c.AbortWithError(http.StatusOK, err)
		return
	}
	// 使用session
	// userInfo := user_response.NewUserInfo(u)
	// middleware.SetSession(userInfo, c)
	userToken := &user.UserToken{Id: u.Id, Username: u.Username, Email: u.Email}
	tokenString, _ := utils.GenerateToken(*userToken)
	loginResponseVo := user_response.NewLoginResponse(*u, tokenString)
	response.OkWithData(loginResponseVo, c)
}

func (l *LoginApi) Logout(c *gin.Context) {
	// 使用session管理登录时要清除
	// if err := middleware.ClearSession(c); err != nil {
	// 	response.Fail(c)
	// 	log.Println(err)
	// } else {
	// 	response.Ok(c)
	// }
	// jwt没有清除操作，可以考虑用redis管理黑名单
	response.OkWithData("退出成功", c)
}
