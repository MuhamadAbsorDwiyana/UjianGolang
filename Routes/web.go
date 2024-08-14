package routes

import (
	ct "UjianGolang/Controllers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func Setup() {
	// Init fiber and fiber engine to Views folder
	engine := html.New("./Views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/css", "./Public/css")
	app.Static("/js", "./Public/js")
	app.Static("/lib", "./Public/lib")
	app.Static("/storage", "./Public/app/storage")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":   "Homepage",
			"AppName": os.Getenv("APP_NAME"),
		}, "layouts/app")
	})

	// Routing stuff with controllers
	app.Get("/visitor", ct.ListVisitor)
	app.Get("/visitor/create", func(c *fiber.Ctx) error {
		return c.Render("dashboard/create", fiber.Map{
			"Title":   "Homepage",
			"AppName": os.Getenv("APP_NAME"),
		}, "layouts/app")
	})
	app.Post("/visitor/create", ct.AddVisitor)
	app.Get("/visitor/update/:id", ct.EditVisitorForm)
	app.Post("/visitor/update/:id", ct.EditVisitor)
	app.Get("/visitor/delete/:id", ct.DeleteVisitorConfirmation)
	app.Post("/visitor/delete/:id", ct.DeleteVisitor)
	app.Get("/visitor/:id", ct.GetVisitor)

	log.Fatal(app.Listen(":3000"))
}
