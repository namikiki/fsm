package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService struct {
	providedKey []byte
	ExpiresAt   time.Duration
}

type customClaims struct {
	UserID string
	jwt.RegisteredClaims
}

func NewJWTService(providedKey []byte, expiresAt time.Duration) *JWTService {
	return &JWTService{providedKey: providedKey, ExpiresAt: expiresAt}
}

// Gen 将传入UUID 存储至JWT内，返回生成的JWT
func (j *JWTService) Gen(ctx context.Context, uuid string) (string, error) {
	if uuid == "" {
		return "", errors.New("uuid 不能为空")
	}

	timeNow := time.Now()
	customClaims := customClaims{
		UserID: uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(timeNow.Add(j.ExpiresAt)),
			IssuedAt:  jwt.NewNumericDate(timeNow),
			NotBefore: jwt.NewNumericDate(timeNow),
		},
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	signedString, err := claims.SignedString(j.providedKey)
	if err != nil {
		return "", fmt.Errorf(" 生成 JWT 时发生错误：%v", err)
	}

	return signedString, nil
}

// Parse 解析传入的JWT，返回UUID
func (j *JWTService) Parse(ctx context.Context, tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.providedKey, nil
	})

	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
		return claims.UserID, nil
	}
	return "", fmt.Errorf("%v", err)
}
