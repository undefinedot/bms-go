package utils

import (
	"bms-go/global"
	"bms-go/model/request"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(global.SYS_CONFIG.JWT.SecretKey),
	}
}

// CreateClaims 根据BaseClaims得到CustomClaims
func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	return request.CustomClaims{
		BaseClaims: baseClaims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(global.SYS_CONFIG.JWT.ExpireTime)).Unix(),
			Issuer:    global.SYS_CONFIG.JWT.Issuer,
		},
	}
}

// CreateToken 生成jwt
func (j *JWT) CreateToken(cliams request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析jwt
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	claims := new(request.CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return claims, nil
	}
	return nil, errors.New("jwt无效")
}
