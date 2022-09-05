package service

// 存放业务逻辑的代码

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

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