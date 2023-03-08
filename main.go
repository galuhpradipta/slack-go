package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

var (
	slackConn    *slack.Client
	slackChannel = ""
)

func createSlackConn(token string) {
	slackToken := os.Getenv("SLACK_TOKEN")
	if slackToken == "" {
		log.Fatal("SLACK_TOKEN is not set")
	}
	slackConn = slack.New(token)
}

func main() {
	slackToken := os.Getenv("SLACK_TOKEN")
	if slackToken == "" {
		log.Fatal("SLACK_TOKEN is not set")
	}

	slackChannel = os.Getenv("SLACK_CHANNEL")
	if slackChannel == "" {
		log.Fatal("SLACK_CHANNEL is not set")
	}
	createSlackConn(slackToken)

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

	if req.Type == "SpamNotification" {
		_, _, _, err := slackConn.SendMessage(slackChannel,
			slack.MsgOptionText(fmt.Sprintf("%s, Email: %s", req.Description, req.Email), false))
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
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
