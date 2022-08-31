package service

// 存放业务逻辑的代码

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) {
	// 1.判断用户是否存在
	mysql.QueryUserByUsername()

	// 2.生成UID
	snowflake.GenID()

	// 3.用户密码加密

	// 4.保存进数据库
	mysql.InsertUser()
}
