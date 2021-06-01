package main

import (
	"log"
	"os"

	"github.com/Haizza1/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	l := log.New(os.Stdout, "api ", log.LstdFlags)
	address := os.Getenv("ADDRESS")
	ph := handlers.NewProducts(l)

	app := fiber.New()
	api := app.Group("/api")
	api.Use(ph.ValidateProduct)

	api.Get("/products", ph.GetProducts)
	api.Post("/products", ph.AddProduct)
	api.Put("/products/:id", ph.UpdateProduct)

	if err := app.Listen(address); err != nil {
		app.Shutdown()
	}
}
