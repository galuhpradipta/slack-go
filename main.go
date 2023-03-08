package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Post("/alert", sendAlert)
	log.Fatal(app.Listen(":3000"))
}

func sendAlert(c *fiber.Ctx) error {
	var req Request
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"body":    req,
	})
}

type Request struct {
	RecordType    string
	Type          string
	TypeCode      int
	Name          string
	Tag           string
	MessageStream string
	Description   string
	Email         string
	From          string
	BouncedAt     time.Time
}
