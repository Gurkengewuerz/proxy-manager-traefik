package routes

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"traefikmanager/server/database"
	"traefikmanager/server/models"
	"traefikmanager/server/traefik"
)

func GenerateConfig(c *fiber.Ctx) error {
	file, err := os.ReadFile("traefik.yml")
	if err != nil {
		return c.Status(500).SendString("500: failed to read config")
	}
	c.Append("Content-Type", "application/x-yaml")
	return c.Send(file)
}

func Commit(c *fiber.Ctx) error {
	err := os.WriteFile("traefik.yml", []byte(traefik.GenerateConfig()), 664)
	if err != nil {
		return c.Status(500).SendString("Failed to generate config")
	}

	database.DBConn.Create(&models.LogEntry{
		User:     "",
		Action:   models.LogActionCommit,
		Metadata: "{}",
	})

	return c.SendString("Ok")
}

func Stats(c *fiber.Ctx) error {
	return c.SendString("fiber")
}
