package handlers

import (
	"log"

	"github.com/Haizza1/api/models"
	"github.com/gofiber/fiber/v2"
)

// Products is a http.Handler
type Products struct {
	log *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(c *fiber.Ctx) error {
	lp := models.GetProducts()
	return c.JSON(lp)
}

func (p *Products) AddProduct(c *fiber.Ctx) error {
	prod := new(models.Product)
	if err := c.BodyParser(prod); err != nil {
		return fiber.NewError(400, "Unable to parse")
	}

	models.AddProduct(prod)
	return c.SendString("created successfully")
}

func (p *Products) UpdateProduct(c *fiber.Ctx) error {
	prod := new(models.Product)
	if err := c.BodyParser(prod); err != nil {
		return fiber.NewError(400, "Unable to parse")
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(400, "Invalid id")
	}

	err = models.UpdateProduct(uint(id), prod)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return c.SendString("Update succesfully")
}
