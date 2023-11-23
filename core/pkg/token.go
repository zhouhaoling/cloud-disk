package pkg

import (
	"cloud/core/define"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken 生成token
func GenerateToken(id int, identity, name string, second int) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	toString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return toString, nil
}

// AnalyzeToken 解析token
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, errors.New("token is invalid")
	}
	return uc, err
}
