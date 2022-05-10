package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"traefikmanager/server/database"
	"traefikmanager/server/models"
)

func GetAudit(c *fiber.Ctx) error {
	var entries []models.LogEntry
	database.DBConn.Order("id desc").Find(&entries)

	entriesTranslated := lo.Map(entries, func(entry models.LogEntry, _ int) models.LogEntry {
		entry.Translated = entry.Translate()
		return entry
	})

	return c.Status(200).JSON(entriesTranslated)
}
