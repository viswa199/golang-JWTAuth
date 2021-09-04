package routes

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Dashboard(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	exp:=fmt.Sprintf("%s",claims.ExpiresAt)
	return ctx.SendString(exp)
}
