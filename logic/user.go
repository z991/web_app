package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) {
	// 判断用户存不存在
	mysql.QueryUserByUsername()
	// 生成 UID
	snowflake.GenID()
	// 数据写入数据库
	mysql.InsertUser()

}
