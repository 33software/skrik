package auth

import (
	"skrik/internal/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error{
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return entities.NewBadRequestError("couldn't verify token. debug: empty token")
		}
		token ,err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			return entities.NewUnauthorizedError("unauthorized. debug: ")
		}


		claims := token.Claims.(jwt.MapClaims)
		useridfloat, ok := claims["userid"].(float64)
		if !ok {
			return entities.NewInternalServerError("internal server error. debug: failed type assertion")
        }
		c.Locals("userid", int(useridfloat))

		
		return c.Next()
	}
}
