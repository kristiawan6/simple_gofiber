package routes

import (
	monthcontroller "fetchAPI_gofiber/src/controllers/MonthController"
	"github.com/gofiber/fiber/v2"
)

// Router configures the routes for the application
func Router(app *fiber.App) {
	v1 := app.Group("/api/v1")

	month := v1.Group("/month")
	month.Get("/data", monthcontroller.GetAllMonths)
}
