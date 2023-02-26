package apis

import (
	"net/http"
	"test/app/services"
	"test/helpers"

	"github.com/gofiber/fiber/v2"
)

func (api *API) HandlerDeleteTargetConfiguration() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			Target string `json:"target"`
		}

		var req request

		if errBody := c.BodyParser(&req); errBody != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   "BodyParser:" + errBody.Error(),
			})
		}

		errDelete := api.serviceMain.DeleteTargetConfiguration(&services.ParamsDeleteTargetConfiguration{
			Target: req.Target,
		})
		if errDelete != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   helpers.ReplEOL(errDelete.Error()),
			})
		}

		c.Status(http.StatusOK)

		return nil
	}
}
