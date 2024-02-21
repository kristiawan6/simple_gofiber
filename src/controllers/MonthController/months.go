package monthcontroller

import (

	"github.com/gofiber/fiber/v2"
	models "fetchAPI_gofiber/src/models/MonthModel"
)

// GetAllMonths returns all months in JSON format
func GetAllMonths(c *fiber.Ctx) error {
	months := models.SelectAllMonth()
	return c.JSON(months)
}
