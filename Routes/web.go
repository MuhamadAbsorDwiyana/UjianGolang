package routes

import (
	ct "UjianGolang/Controllers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func Setup() {
	// Init fiber and fiber engine to Views folder
	engine := html.New("./Views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Routing stuff with controllers
	app.Get("/", ct.ListVisitor)
	app.Get("/visitor", func(c *fiber.Ctx) error {
		return c.Render("create", nil)
	})
	app.Post("/visitor", ct.AddVisitor)

	log.Fatal(app.Listen(":3000"))
}
