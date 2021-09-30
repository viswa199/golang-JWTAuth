package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logout(ctx *fiber.Ctx) error {
	cookie:=&fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	ctx.Cookie(cookie)
	return ctx.Redirect("/login")
}