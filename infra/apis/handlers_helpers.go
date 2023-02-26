package apis

import (
	"github.com/gofiber/fiber/v2"
)

func (api *API) HandlerHealth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("alive")
	}
}
