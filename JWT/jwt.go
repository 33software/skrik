package jwtGen

import (
	"audio-stream-golang/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(userid uint) (string, error) {
	EnvConfig := config.GetConfig()
	claims := jwt.MapClaims{
		"userid": userid,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return newToken.SignedString(EnvConfig.Jwt_keyword)
}
