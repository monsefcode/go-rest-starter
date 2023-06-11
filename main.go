package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/monsefcode/go-rest-starter/database"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	err := initApp()

	if err != nil {
		panic(err)
	}
	// close db connection
	defer database.CloseMongoDB()

	app := fiber.New()

	app.Use(logger.New())

    app.Post("/", func(c *fiber.Ctx) error {
		sampleDoc := bson.M{"name": "John Doe"}
		collection := database.GetCollection("users")

		newDoc, err := collection.InsertOne(context.TODO(), sampleDoc)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error while inserting document")
		}

		// sending the new document
		return c.JSON(newDoc)
    });

	// testing route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

    app.Listen(":8000")
}


// initApp initializes the application
func initApp() error {
	err := loadENV()
	if err != nil {
		return err
	}

	err = database.StartMongoDB()
	if err != nil {
		return err
	}

	return nil
}

func loadENV() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file | loadENV() - main.go")
		return err
	}
	return nil
}