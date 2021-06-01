package handlers

import (
	"github.com/Haizza1/api/models"
	"github.com/gofiber/fiber/v2"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// reponses:
//	200: productsResponse
//	404: noProductsResponse

// GetProducts returns the data from the data store
func (p *Products) GetALl(c *fiber.Ctx) error {
	lp := models.GetProducts()
	return c.JSON(lp)
}
