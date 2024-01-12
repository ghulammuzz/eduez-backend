// routes/product.go

package routes

import (
	"native/crus/internal/handlers/guide"

	"github.com/gofiber/fiber/v2"
)

func SetupGuideRoutes(app *fiber.App) {
	guideGroup := app.Group("/guide")
	// guideGroup.Get("/", guide.ListGuide)
	guideGroup.Get("/", guide.ListGuideWithPagination)
	guideGroup.Get("/me", guide.ListMyGuide)
	guideGroup.Post("/", guide.CreateGuide)
	guideGroup.Get("/:id", guide.DetailGuide)
	guideGroup.Put("/mark", guide.MarkCompleteSubtitle)
}
