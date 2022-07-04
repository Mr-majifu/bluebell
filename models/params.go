package models

type ParamRegister struct {
	UserName   string `json:"user_name" binding:"required"`
	Password   string `json:"password" binding:"required,eqcsfield=RePassword"`
	RePassword string `json:"re_password" binding:"required"`
}

type ParamLogin struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
