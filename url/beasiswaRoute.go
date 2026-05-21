package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

func BeasiswaRoute(app *fiber.App) {
	app.Get("/beasiswa", controller.GetBeasiswa)
	app.Get("/beasiswa/:id", controller.GetDetailBeasiswa)
	app.Post("/beasiswa", controller.AddBeasiswa)
	app.Put("/beasiswa/:id", controller.UpdateBeasiswa)
	app.Delete("/beasiswa/:id", controller.DeleteBeasiswa)
}
