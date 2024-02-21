package monthcontroller

import (
	"encoding/json"
	models "fetchAPI_gofiber/src/models/MonthModel"

	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetAllMonths(c *fiber.Ctx) error {
	months := models.SelectAllMonth()
	res, err := json.Marshal(months)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Gagal Konversi Json")
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(res)
}
