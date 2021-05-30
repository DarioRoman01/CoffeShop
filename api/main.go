package main

import (
	"log"
	"os"

	"github.com/Haizza1/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	l := log.New(os.Stdout, "api ", log.LstdFlags)
	port := os.Getenv("PORT")
	ph := handlers.NewProducts(l)

	app := fiber.New()
	api := app.Group("/api")

	api.Get("/products/:id", ph.GetProducts)
	api.Post("/products", ph.AddProduct)
	api.Put("/products/:id", ph.UpdateProduct)

	if err := app.Listen(port); err != nil {
		app.Shutdown()
	}
}
