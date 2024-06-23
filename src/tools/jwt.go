package tools

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

func init() {
	jwtSecret = []byte(fmt.Sprintf("xx_xx_xx_xx_xx_xx_xx_%s", time.Now().Format(Short)))
}

// Claims 载体
type Claims struct {
	Who string `json:"who"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(who string) (string, error) {
	claims := Claims{
		MD5(who),
		jwt.StandardClaims{
			Subject:   "4login",                                           //主题
			ExpiresAt: time.Now().Add(time.Hour * 24 * 5).Unix(),          //过期时间
			Issuer:    fmt.Sprintf("issuer_%s", time.Now().Format(Short)), //签发人
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	if err != nil {
		//https://gowalker.org/github.com/dgrijalva/jwt-go#ValidationError
		//jwt.ValidationError 是一个无效token的错误结构
		if ve, ok := err.(*jwt.ValidationError); ok {
			// ValidationErrorMalformed是一个uint常量，表示token不可用
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("TOKEN 不可用")
				// ValidationErrorExpired表示Token过期
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("TOKEN 已过期")
				// ValidationErrorNotValidYet表示无效token
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, fmt.Errorf("无效的 TOKEN")
			} else {
				return nil, fmt.Errorf("TOKEN 不可用")
			}
		}
	}

	return nil, err
}
