package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"github.com/sizata-siege/finance-management/routes"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	/* Load .env */
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	/* Connect to DB using kamva/mgm (https://github.com/Kamva/mgm) */
	if err := mgm.SetDefaultConfig(
		nil,
		"finance",
		options.Client().ApplyURI("mongodb://localhost:27017"),
	); err != nil {
		log.Fatal("Error connecting to db: ", err)
	}
	log.Println("Connected to db ✔️")
}

func main() {
	log.Println("Starting Server")
	app := fiber.New(fiber.Config{
		AppName: "Sizata's Finance Management System",
		// Prefork: true,
	})

	app.Use(logger.New())
	app.Use(recover.New())

	routes.SetupAPI(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8070"
	}
	log.Println("Listening on port", port)
	log.Fatal(app.Listen(":" + port))
}
