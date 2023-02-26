package rest

import (
	"errors"
	"net/http"

	"test/app/apperrors"
	"test/helpers"

	"github.com/gofiber/fiber/v2"
)

func (rest *Rest) HandlerNewItem() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			Email string `json:"email"`
		}

		var req request

		if errBody := c.BodyParser(&req); errBody != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   "BodyParser:" + errBody.Error(),
			})
		}

		var itemID int
		var errIns error

		// itemID, errIns := CreateItem(c.Context(), &ParamsCreateItem{
		// 	Email: req.Email,
		// })
		if errIns != nil {
			if errors.As(errIns, &apperrors.ErrValidation{}) {
				return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
					"success": false,
					"error":   helpers.ReplEOL(errIns.Error()),
				})
			}

			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   helpers.ReplEOL(errIns.Error()),
			})
		}

		return c.Status(http.StatusOK).JSON(&fiber.Map{
			"success": true,
			"id":      itemID,
		})
	}
}
