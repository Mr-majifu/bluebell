package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

func Register(p *models.ParamRegister) (err error) {
	// 判断用户是否存在
	if err = mysql.CheckUserExist(p.UserName); err != nil {
		return err
	}
	// 生成UID
	uid := snowflake.GenID()
	// 构造User实例
	user := &models.User{
		UserID:   uid,
		UserName: p.UserName,
		Password: p.Password,
	}
	// 保存进数据库
	if err = mysql.InsertUser(user); err != nil {
		return err
	}
	return
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID了
	if err = mysql.Login(user); err != nil {
		return "", err
	}
	// 生成JWT的Token
	return jwt.GenToken(user.UserID, user.UserName)
}
