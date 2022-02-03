package utils

import (
	"errors"
	"log"
	"paper-manager/model/user"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	User user.UserToken
	jwt.StandardClaims
}

const (
	TokenExpireDuration       = time.Hour * 2
	TokenExpireDurationMinute = time.Minute * 5
)

var MySecret = []byte("sacadfqwda")

func GenerateToken(userInfo user.UserToken) (string, error) {
	expirationTime := time.Now().Add(TokenExpireDuration)
	// expirationTime := time.Now().Add(TokenExpireDurationMinute)
	claims := &MyClaims{
		User: userInfo,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "syf",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err := token.SignedString(MySecret); err != nil {
		log.Println("Token签名出错")
		return "", err
	} else {
		return tokenString, nil
	}

}

func RenewToken(claims *MyClaims) (string, error) {
	// expiredTime := time.Unix(cliams.ExpiresAt, 0)
	// 若token过期不超过10分钟则给它续签
	if WithinLimit(claims.ExpiresAt, 600) {
		return GenerateToken(claims.User)
	}
	return "", errors.New("登录已过期")
}

func ParseToken(tokenString string) (*MyClaims, error) {
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		if strings.Contains(err.Error(), "expired") {
			return claims, err
		}
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Token 校验出错")
	}
	return claims, nil

}
