package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户存不存在

	if err := mysql.CheckUserExists(p.Username); err != nil {
		return err
	}

	// 生成 UID
	userID := snowflake.GenID()
	// 构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 数据写入数据库
	return mysql.InsertUser(user)

}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	// 传递user指针，获取User.UserID
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	// 生成 JWT token
	return jwt.GenToken(user.UserID, user.Username)

}
