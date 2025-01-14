package auth

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error{
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "empty token"})
		}
		token ,err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			log.Println("invalid token!")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}


		claims := token.Claims.(jwt.MapClaims)
		useridfloat, ok := claims["userid"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "invalid userid in token",
            })
        }
		c.Locals("userid", int(useridfloat))

		
		return c.Next()
	}
}
