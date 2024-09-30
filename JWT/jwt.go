package jwtGen

import (
	"audio-stream-golang/config"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(userid uint) (string, error) {
	EnvConfig := config.GetConfig()
	claims := jwt.MapClaims{
		"userid": userid,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return newToken.SignedString([]byte(EnvConfig.Jwt_keyword))
}

func JwtProtected() fiber.Handler {
	EnvConfig := config.GetConfig()
	return jwtware.New(jwtware.Config{
		SigningKey:  []byte(EnvConfig.Jwt_keyword),
		AuthScheme:  "Bearer",
		TokenLookup: "header:Authorization",
	})

}
