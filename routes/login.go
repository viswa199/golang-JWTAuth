package routes

import (
	"JWTAuth/database"
	"JWTAuth/hash"
	"JWTAuth/users"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var Credentials users.Autheticate

func LoginPage(ctx *fiber.Ctx) error {
	return ctx.Render("templates/login.gohtml", nil)
}

func Login(ctx *fiber.Ctx) error {
	//working with database
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	database.Client, err = database.Connect()
	if err != nil {
		return fiber.DefaultErrorHandler(ctx, err)
	}
	database.Client.Connect(c)
	defer database.Client.Disconnect(c)

	//checking user credentials with credentials stored in database.
	Credentials.Username = ctx.FormValue("username")
	Credentials.Password = ctx.FormValue("password")
	GoDatabase := database.Client.Database("golang")
	GoCollection := GoDatabase.Collection("users")
	var body bson.M
	err = GoCollection.FindOne(c, bson.D{
		{Key: "Username", Value: Credentials.Username},
	}).Decode(&body)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.DefaultErrorHandler(ctx, err)
		} else {
			return fiber.DefaultErrorHandler(ctx, err)
		}
	}
	Password := fmt.Sprintf("%s", body["Password"])
	if hash.DecodePassword(Password, Credentials.Password) {
		token, err := GenerateToken()
		if err != nil {
			return fiber.DefaultErrorHandler(ctx, err)
		}
		cookie:=fiber.Cookie{
			Name: "jwt",
			Value: token,
			Expires: time.Now().Add(time.Minute*30),
			HTTPOnly: true,
		}
		ctx.Cookie(&cookie)
		return ctx.JSON(fiber.Map{
			"message": "Login Succeeded",
		})
	}
	return ctx.SendString("Invalid Username and Password.")
}
