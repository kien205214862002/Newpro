package utils

import (
	"errors"
	"fmt"
	"go01-airbnb/config"
	"go01-airbnb/pkg/common"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Data nhận được sau khi tạo token
type Token struct {
	AccessToken string `json:"accessToken"`
	ExpiresAt   int64  `json:"expiresAt"`
}

// Data truyền vào để tạo token
type TokenPayload struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type myClaims struct {
	Payload TokenPayload `json:"payload"`
	jwt.RegisteredClaims
}

// type jwtProvider struct {
// 	secret string
// }
// func NewJWTProvider(secret string) *jwtProvider {
// 	return &jwtProvider{secret}
// }

func GenerateJWT(data TokenPayload, cfg *config.Config) (*Token, error) {
	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Hour * 12))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.RegisteredClaims{
			// Token sẽ hết hạn khi nào
			ExpiresAt: expiresAt,
			// Token được tạo khi nào
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ID:       fmt.Sprintf("%d", time.Now().UnixNano()),
		},
	})

	accessToken, err := token.SignedString([]byte(cfg.App.Secret))
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken: accessToken,
		ExpiresAt:   expiresAt.Unix(),
	}, nil
}

func ValidateJWT(accessToken string, cfg *config.Config) (*TokenPayload, error) {
	token, err := jwt.ParseWithClaims(accessToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.App.Secret), nil
	})

	if err != nil {
		return nil, ErrInvalidToken
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*myClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return &claims.Payload, nil
}

// Khai báo lỗi liên quan đến token
var (
	ErrTokenNotFound = common.ErrUnauthorized(
		errors.New("token not found"),
	)

	ErrEncodingToken = common.ErrUnauthorized(
		errors.New("error encoding token"),
	)

	ErrInvalidToken = common.ErrUnauthorized(
		errors.New("invalid token"),
	)
)
