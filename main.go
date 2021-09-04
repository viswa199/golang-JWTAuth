package main

import (
	"JWTAuth/database"
	"JWTAuth/routes"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

var err error

func init() {
	//connecting to database
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	database.Client, err = database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	database.Client.Connect(c)
	defer database.Client.Disconnect(c)
}

func main() {
	//creating new fiber instance
	app := fiber.New(fiber.Config{
		Prefork:       true,
		StrictRouting: true,
	})

	//setting up our roots
	app.Get("/", Home)
	app.Post("/register", routes.Register)
	app.Get("/loginPage", routes.LoginPage)
	app.Post("/login", routes.Login)
	app.Get("/dashboard",routes.Dashboard)

	//setting up server
	app.Listen(":8080")
}

func Home(ctx *fiber.Ctx) error {
	return ctx.Render("templates/index.gohtml", nil)
}
