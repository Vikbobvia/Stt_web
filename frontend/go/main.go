package main

import (
	// "fmt"
	"log"

	"github.com/gofiber/fiber/v2"

)

func main() {
	app := fiber.New()

	app.Get("/", func( ctx * fiber.Ctx) error {
		return ctx.Status(200).JSON(map[string]interface{}{"msg":"hello world"})
	})

	log.Fatal(app.Listen(":8080"))
}
