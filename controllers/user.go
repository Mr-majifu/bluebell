package controllers

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	// 1、获取参数和参数校验
	var p = new(models.ParamRegister)
	if err := c.ShouldBind(p); err != nil {
		// 记录日志
		zap.L().Error("Register with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 手动对请求参数进行详细的业务逻辑判断
	//if len(p.UserName) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	//	zap.L().Error("Register with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}

	// 2、业务逻辑处理
	if err := logic.Register(p); err != nil {
		zap.L().Error("logic.Register failed:", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3、正常响应
	ResponseSuccess(c, nil)
}

func Login(c *gin.Context) {
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Printf("ctx.Request.body: %v", string(data))
	// 获取参数
	var p = new(models.ParamLogin)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("Login with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 业务逻辑处理
	token, err := logic.Login(p)
	if err != nil {
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}

	// 正常响应
	ResponseSuccess(c, token)
}
