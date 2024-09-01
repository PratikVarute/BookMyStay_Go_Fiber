package main

import (
	"context"
	"flag"
	"log"

	"github.com/PratikVarute/BookMyStay_Go_Fiber/api"
	"github.com/PratikVarute/BookMyStay_Go_Fiber/db"
	"github.com/PratikVarute/BookMyStay_Go_Fiber/literals"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {

	listenAddr := flag.String("listenAddr", ":5000", "default local server addr..")
	app := fiber.New(config)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(literals.DbUrl))
	if err != nil {
		log.Fatal(err)
	}
	//handler intilization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	apiV1 := app.Group("/api/v1")
	apiV1.Get("/user/:id", userHandler.HandelGetUser)
	apiV1.Get("/users", userHandler.HandelGetUsers)
	apiV1.Post("/new-user", userHandler.HandleInsertUser)

	app.Listen(*listenAddr)
}
