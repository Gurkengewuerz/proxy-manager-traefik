package routes

import (
	"github.com/gofiber/fiber/v2"
	"traefikmanager/server/claims"
)

func AuthInfoMiddleware(c *fiber.Ctx) error {
	userClaims := c.Locals("claims").(*claims.IDTokenClaims)
	return c.Status(200).JSON(userClaims)
}
