package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpirDuration = time.Hour * 2

var mySecret = []byte("这是我的密码")

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// GenToken 生成token
func GenToken(userID int64, userName string) (token string, err error) {
	c := MyClaims{
		userID,
		userName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpirDuration).Unix(), // 过期时间
			Issuer:    "bluebell",                                // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象,尼玛，是HS256，不是ES256
	// 使用指定的secret签名并获得完整的编码后的字符串token（加盐？）
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)
	return
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims) // 这里必须要手动初始化 return 的值！！
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

/// GenToken 生成token
//func GenToken(userID int64, userName string) (aToken, rToken string, err error) {
//	c := MyClaims{
//		userID,
//		userName,
//		jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(TokenExpirDuration).Unix(), // 过期时间
//			Issuer:    "bluebell",                                // 签发人
//		},
//	}
//	// 使用指定的签名方法创建签名对象,尼玛，是HS256，不是ES256
//	// 使用指定的secret签名并获得完整的编码后的字符串token（加盐？）
//	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)
//	// refresh token 不需要加自定义数据
//	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
//		ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
//		Issuer:    "bluebell",
//	}).SignedString(mySecret)
//	return
//}

//token
//func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
//	if _, err = jwt.Parse(rToken,Keyfunc)
//}
