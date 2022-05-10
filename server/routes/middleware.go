package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"traefikmanager/server/claims"
	"traefikmanager/server/database"
	models2 "traefikmanager/server/models"
)

func GetMiddleware(c *fiber.Ctx) error {
	var middlewares []models2.Middleware
	database.DBConn.Preload("Settings").Find(&middlewares)

	return c.Status(200).JSON(middlewares)
}

func PutMiddleware(c *fiber.Ctx) error {
	return c.SendString("fiber")
}

func PostMiddleware(c *fiber.Ctx) error {
	jsonMiddleware := new(models2.Middleware)
	if err := c.BodyParser(jsonMiddleware); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if !jsonMiddleware.Validate() {
		return c.Status(400).JSON("failed to validate input data")
	}

	database.DBConn.Create(&jsonMiddleware)

	userClaims := c.Locals("claims").(*claims.IDTokenClaims)
	jsonData, jsonErr := json.Marshal(map[string]string{
		"type": string(rune(jsonMiddleware.Type)),
	})
	if jsonErr == nil {
		database.DBConn.Create(&models2.LogEntry{
			User:     userClaims.Username,
			Action:   models2.LogActionCreateMiddleware,
			Metadata: string(jsonData),
		})
	}

	return c.Status(201).JSON("ok")
}

func DeleteMiddleware(c *fiber.Ctx) error {
	var middlewares []models2.Middleware
	jsonMiddleware := new(models2.Middleware)
	if err := c.BodyParser(jsonMiddleware); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DBConn.Where("id = ?", jsonMiddleware.ID).Delete(&middlewares)

	userClaims := c.Locals("claims").(*claims.IDTokenClaims)
	jsonData, jsonErr := json.Marshal(map[string]string{
		"id":   strconv.Itoa(int(jsonMiddleware.ID)),
		"type": string(rune(jsonMiddleware.Type)),
	})
	if jsonErr == nil {
		database.DBConn.Create(&models2.LogEntry{
			User:     userClaims.Username,
			Action:   models2.LogActionDeleteMiddleware,
			Metadata: string(jsonData),
		})
	}

	return c.Status(200).JSON("deleted")
}
