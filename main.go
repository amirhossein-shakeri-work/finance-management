package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kamva/mgm/v3"
	"github.com/sizata-siege/finance-management/account"
	"github.com/sizata-siege/finance-management/session"
	"github.com/sizata-siege/finance-management/user"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupRoutes(app *fiber.App) {
	// app.Get("/", indexHome)
	app.Static("/", "./public", fiber.Static{MaxAge: -1})

	api := app.Group("/api", callNext)

	v1 := api.Group("/v1", callNext)
	/* =-=-=-=-=-=-= Accounts =-=-=-=-=-=-= */
	v1.Get("/accounts", account.Index)
	v1.Post("/accounts", account.Store)
	v1.Get("/accounts/:id", account.Show)
	v1.Patch("/accounts/:id", account.Update)
	v1.Delete("/accounts/:id", account.Delete)
	/* =-=-=-=-=-=-= Session & User =-=-=-=-=-=-= */
	// v1.Get("/session")    // get loged in user
	v1.Post("/session", session.Login)   // login
	v1.Delete("/session", session.Logout) // logout / smiliar to /logout
	v1.Post("/users", user.CreateNewUser)
	/* =-=-=-=-=-=-= Transactions =-=-=-=-=-=-= */
}

func init() {
	// https://github.com/Kamva/mgm
	err := mgm.SetDefaultConfig(nil, "finance",
		options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to db ...")
}

func main() {
	log.Println("Starting Server")
	app := fiber.New(fiber.Config{
		AppName: "Sizata's Finance Management System",
		// Prefork: true,
	})

	app.Use(logger.New())
	app.Use(recover.New())

	setupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8070"
	}
	log.Println("Listening on port", port)
	log.Fatal(app.Listen(":" + port))
}

func indexHome(c *fiber.Ctx) error {
	return c.SendString("Welcome to SIZATA's Finance Management System")
}

func callNext(c *fiber.Ctx) error { return c.Next() }
