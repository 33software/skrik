package auth

import (
	"skrik/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateAccessToken(id uint) (string, error){
    claims := jwt.MapClaims{
        "userid": id,
        "exp": time.Now().Add(time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(config.AppConfig.Jwt_keyword))
}
func ParseToken (tokenString string) (*jwt.Token, error){
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
        return []byte(config.AppConfig.Jwt_keyword), nil
    })
}