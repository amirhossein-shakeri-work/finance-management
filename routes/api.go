package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sizata-siege/finance-management/auth"
	"github.com/sizata-siege/finance-management/controllers"
)

func SetupAPI(app *fiber.App) {
	app.Get("/", indexHome)
	app.Static("/", "./public", fiber.Static{MaxAge: 0})

	api := app.Group("/api", callNext)

	v1 := api.Group("/v1", callNext)

	/* =-=-=-=-=-=-= Accounts =-=-=-=-=-=-= */
	accounts := v1.Group("/accounts", auth.Middleware)
	accounts.Get("/", controllers.IndexAccounts)
	accounts.Post("/", controllers.CreateAccount)
	accounts.Get("/:id", controllers.ShowAccount)
	accounts.Patch("/:id", controllers.UpdateAccount)
	accounts.Delete("/:id", controllers.DeleteAccount)

	/* =-=-=-=-=-=-= Session & User =-=-=-=-=-=-= */
	// v1.Get("/auth")    // get loged in user
	v1.Post("/session", auth.Login)                     // login
	v1.Delete("/session", auth.Middleware, auth.Logout) // logout / smiliar to /logout
	v1.Post("/users", auth.CreateNewUser)

	/* =-=-=-=-=-=-= Transactions =-=-=-=-=-=-= */
	transactions := v1.Group("/transactions", auth.Middleware)
	transactions.Get("/", controllers.IndexTransactions)
	transactions.Post("/", controllers.CreateTransaction)

	/* =-=-=-=-=-=-= Test =-=-=-=-=-=-= */
	app.Get("/test", auth.Middleware, testHandler)
}

const welcomeMessage = "Welcome to SIZATA's Finance Management System"

func indexHome(c *fiber.Ctx) error {
	/* if request doesn't accepts html */
	if c.Accepts("text/html") == "" {
		return c.JSON(fiber.Map{"message": welcomeMessage})
	}
	return c.Next()
}

func callNext(c *fiber.Ctx) error { return c.Next() }

func testHandler(c *fiber.Ctx) error {
	// c.Cookie(&fiber.Cookie{
	// 	Name: "Foo",
	// 	Value: "Bar",
	// 	Expires: time.Now().Add(time.Minute * 10),
	// })

	// log.Printf("%v", usr.Claims)
	// j := c.Locals("user").(*jwt.Token)
	// claims := j.Claims.(jwt.MapClaims)
	// fmt.Printf("%v %v !!! %T %T\n", claims["id"], claims["exp"], claims["exp"], claims["id"])
	// j := jwt.New(c)
	// fmt.Println(j.User, j.Claims, j.User.ID)
	// u := &user.User{}
	// er := mgm.Coll(u).First(bson.M{"email":"maiddl@example.com"}, u)
	// fmt.Println(*u, u == nil, u.Email == "", er)
	return c.SendString("OK")
}
