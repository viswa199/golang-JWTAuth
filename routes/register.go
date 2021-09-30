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
)

var User users.User
var err error 

func Register(ctx *fiber.Ctx) error {
	database.Client,err=database.Connect()
	if err!=nil{
		return fiber.DefaultErrorHandler(ctx,err)
	}
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	database.Client.Connect(c)
	defer database.Client.Disconnect(c)
	User.Username = ctx.FormValue("username")
	User.Email = ctx.FormValue("email")
	User.Password = hash.HashPassword(ctx.FormValue("password"))
	GoDatabase := database.Client.Database("golang")
	GoCollection := GoDatabase.Collection("users")
	res, err := GoCollection.InsertOne(c, bson.D{
		{Key: "Username", Value: User.Username},
		{Key: "Email", Value: User.Email},
		{Key: "Password", Value: User.Password},
	})
	if err != nil {
		return fiber.DefaultErrorHandler(ctx, err)
	}
	fmt.Println(res.InsertedID)
	return ctx.Redirect("/login")
}