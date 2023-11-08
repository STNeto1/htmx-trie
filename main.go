package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stneto1/trie/pkg"
)

func main() {

	tmpl := pkg.CreateTemplate()
	container := pkg.NewContainer(tmpl)

	app := fiber.New(fiber.Config{
		Views: tmpl,
	})
	// app.Use(logger.New())

	app.Get("/", container.HandleIndex)

	app.Post("/create", container.HandleCreateWord)
	app.Delete("/delete/:index", container.HandleDeleteWord)
	// search word

	app.Listen(":3000")
}
