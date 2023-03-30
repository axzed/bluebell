package service

// 存放业务逻辑的代码

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

// SignUp 处理用户注册业务逻辑
func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		// 数据库查询错误
		return err
	}

	// 2.生成UID
	userID := snowflake.GenID()
	// 构造一个User实例
	user := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 3.保存进数据库(其中对密码进行加密)
	return mysql.InsertUser(&user)
}

// Login 用户登录
func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
