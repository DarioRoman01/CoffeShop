package handlers

import (
	"log"
	"net/http"

	"github.com/Haizza1/api/models"
	"github.com/gofiber/fiber/v2"
)

var (
	parseError = fiber.NewError(400, "Unable to unmarshal json")
	idError    = fiber.NewError(400, "Invalid Id")
)

const keyProduct = "prod"

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
	prod := c.Context().UserValue(keyProduct).(*models.Product)
	models.AddProduct(prod)
	return c.SendString("created successfully")
}

func (p *Products) UpdateProduct(c *fiber.Ctx) error {
	prod := c.Context().UserValue(keyProduct).(*models.Product)
	id, err := c.ParamsInt("id")
	if err != nil {
		return idError
	}

	err = models.UpdateProduct(uint(id), prod)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return c.SendString("Update succesfully")
}

func (p *Products) ValidateProduct(c *fiber.Ctx) error {
	if c.Method() != http.MethodGet {
		prod := new(models.Product)
		if err := c.BodyParser(prod); err != nil {
			return parseError
		}

		if err := prod.Validate(); err != nil {
			return fiber.NewError(400, err.Error())
		}

		c.Context().SetUserValue(keyProduct, prod)
		return c.Next()
	}

	return c.Next()
}
