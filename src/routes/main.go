package routes

import (
	monthcontroller "fetchAPI_gofiber/src/controllers/MonthController"

	"github.com/gofiber/fiber/v2"
)

func Router(c *fiber.App) {

	 v1 := c.Group("/api/v1")

	month := v1.Group("/month")
	{
		month.Get("/data", monthcontroller.GetAllMonths)

	}

}
