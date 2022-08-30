package utils

import (
	"errors"
	"opengin/server/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	ClientId string `json:"client_id"`
	jwt.RegisteredClaims
}

func CreateToken(userId int) (string, error) {
	claims := CustomClaims{
		GetUUID(),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(config.Settings.JWT.ExpireHours))),
			Subject:   strconv.Itoa(userId),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Settings.JWT.SecretKey))
}

func ParseToken(token string) (*CustomClaims, error) {
	payload, err := jwt.ParseWithClaims(
		token,
		&CustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(config.Settings.JWT.SecretKey), nil
		})
	if err != nil {
		return nil, err
	}

	if claims, isOk := payload.Claims.(*CustomClaims); isOk && payload.Valid {
		return claims, nil
	}

	return nil, errors.New("failed to parse this token")
}
