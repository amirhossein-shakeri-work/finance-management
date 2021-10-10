package user

import "github.com/gofiber/fiber/v2"

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func deleteAccount(c *fiber.Ctx) error {
	return nil
}
