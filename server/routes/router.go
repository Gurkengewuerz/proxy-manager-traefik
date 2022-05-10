package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"traefikmanager/server/claims"
	"traefikmanager/server/database"
	models2 "traefikmanager/server/models"
)

func GetRouter(c *fiber.Ctx) error {
	var routers []models2.Router
	database.DBConn.Preload("Locations").Preload("Locations.Middlewares").Preload("Locations.Middlewares.Settings").Find(&routers)

	return c.Status(200).JSON(routers)
}

func PutRouter(c *fiber.Ctx) error {
	return c.SendString("fiber")
}

func PostRouter(c *fiber.Ctx) error {
	jsonRouter := new(models2.Router)
	if err := c.BodyParser(jsonRouter); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if !jsonRouter.Validate() {
		return c.Status(400).JSON("failed to validate input data")
	}

	database.DBConn.Create(&jsonRouter)

	userClaims := c.Locals("claims").(*claims.IDTokenClaims)
	jsonData, jsonErr := json.Marshal(map[string]string{
		"name": jsonRouter.Name,
	})
	if jsonErr == nil {
		database.DBConn.Create(&models2.LogEntry{
			User:     userClaims.Username,
			Action:   models2.LogActionCreateRouter,
			Metadata: string(jsonData),
		})
	}

	return c.Status(201).JSON("ok")
}

func DeleteRouter(c *fiber.Ctx) error {
	var routers []models2.Router
	jsonRouter := new(models2.Router)
	if err := c.BodyParser(jsonRouter); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DBConn.Where("id = ?", jsonRouter.ID).Delete(&routers)

	userClaims := c.Locals("claims").(*claims.IDTokenClaims)
	jsonData, jsonErr := json.Marshal(map[string]string{
		"name": jsonRouter.Name,
	})
	if jsonErr == nil {
		database.DBConn.Create(&models2.LogEntry{
			User:     userClaims.Username,
			Action:   models2.LogActionDeleteRouter,
			Metadata: string(jsonData),
		})
	}

	return c.Status(200).JSON("deleted")
}
