package apis

import (
	"errors"
	"net/http"
	"test/app/apperrors"
	"test/app/services"
	"test/helpers"

	"github.com/gofiber/fiber/v2"
)

func (api *API) HandlerGetCurrentTarget() fiber.Handler {
	return func(c *fiber.Ctx) error {
		content, errGet := api.controller.GetCurrentTarget()
		if errGet != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   helpers.ReplEOL(errGet.Error()),
			})
		}

		return c.Status(http.StatusOK).Send(content)
	}
}

func (api *API) HandlerGetTargets() fiber.Handler {
	return func(c *fiber.Ctx) error {
		content, errGet := api.controller.GetTargets()
		if errGet != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   helpers.ReplEOL(errGet.Error()),
			})
		}

		return c.Status(http.StatusOK).Send(content)
	}
}

func (api *API) HandlerGetVariableValues() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			VName   string   `json:"name"`
			Targets []string `json:"targets"`
		}

		var req request

		if errBody := c.BodyParser(&req); errBody != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   "BodyParser:" + errBody.Error(),
			})
		}

		content, errFind := api.controller.GetVariableValues(&services.ParamsGetVariableValues{
			Name:    req.VName,
			Targets: req.Targets,
		})
		if errFind != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   helpers.ReplEOL(errFind.Error()),
			})
		}

		return c.Status(http.StatusOK).Send(content)
	}
}

func (api *API) HandlerGetTargetConfiguration() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			Target string `json:"target"`
		}

		var req request

		if errBody := c.BodyParser(&req); errBody != nil && errBody.Error() != "unexpected end of JSON input" {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   "BodyParser:" + errBody.Error(),
			})
		}

		content, errFind := api.controller.GetTargetConfiguration(&services.ParamsFindTargetConfiguration{
			Target: req.Target,
		})
		if errFind != nil {
			if errors.As(errFind, &apperrors.ErrValidation{}) {
				return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
					"success": false,
					"error":   helpers.ReplEOL(errFind.Error()),
				})
			}

			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   helpers.ReplEOL(errFind.Error()),
			})
		}

		return c.Status(http.StatusOK).Send(content)
	}
}
