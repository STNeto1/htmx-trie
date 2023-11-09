package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stneto1/trie/pkg"
)

func main() {

	tmpl := pkg.CreateTemplate()
	container := pkg.NewContainer()

	app := fiber.New(fiber.Config{
		Views: tmpl,
	})
	// app.Use(logger.New())
	app.Static("/public", "./public")

	app.Get("/", container.HandleIndex)

	app.Post("/create", container.HandleCreateWord)
	app.Post("/search", container.HandleSearchWord)
	app.Delete("/delete/:index", container.HandleDeleteWord)

	app.Listen(":3000")
}
