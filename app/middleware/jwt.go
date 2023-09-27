package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"os"
)

func JWTConfig(app *fiber.App) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(secretKey),
	}))
}
