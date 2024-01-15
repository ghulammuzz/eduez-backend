package main

import (
	"log"
	"native/crus/config/db"
	"native/crus/internal/routes"

	"github.com/gofiber/fiber/v2"
)

// handler = interface
// mengikuti kontrak (w http.ResponseWriter, r *http.Request)
// kontrak request harus pointern

func main() {
	db.InitDB()
	defer db.CloseDB()

	app := fiber.New(fiber.Config{
		AppName: "EduZe",
		// Prefork: true,
	})
	routes.SetupGuideRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
