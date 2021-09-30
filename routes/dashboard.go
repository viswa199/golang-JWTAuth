package routes

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var (
	Golang_votes int
	Python_votes int
	Java_votes   int
	nodejs_votes int
)

func Dashboard(ctx *fiber.Ctx) error {
	if ctx.Method() == "GET" {
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
		return ctx.Render("templates/dashboard.gohtml", claims)
	}
	if ctx.Method() == "POST" {
		user_vote := ctx.FormValue("fav-language")
		print(UpdateVoting(user_vote))
		return ctx.Redirect("/dashboard")
	}
	return nil
}

func UpdateVoting(user_vote string) int {
	if user_vote == "golang" {
		Golang_votes += 1
		return Golang_votes
	} else if user_vote == "python" {
		Python_votes += 1
		return Python_votes
	} else if user_vote == "java" {
		Java_votes += 1
		return Java_votes
	} else if user_vote == "nodejs" {
		nodejs_votes += 1
		return nodejs_votes
	} else {
		return 0
	}
}
