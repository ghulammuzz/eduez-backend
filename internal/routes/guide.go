// routes/product.go

package routes

import (
	"native/crus/internal/handlers/guide"

	"native/crus/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupGuideRoutes(app *fiber.App) {
	guideGroup := app.Group("/guide", middleware.FirebaseAuth())
	// guideGroup.Get("/", guide.ListGuide)
	guideGroup.Get("/test", guide.Test)
	guideGroup.Get("/", guide.ListGuideWithPagination)
	// guideGroup.Get("/me", guide.ListMyGuide)
	guideGroup.Get("/me", guide.PaginateMyGuide)
	guideGroup.Post("/", guide.CreateGuide)
	guideGroup.Get("/:id", guide.DetailGuide)
	guideGroup.Put("/mark", guide.MarkCompleteSubtitle)
}
