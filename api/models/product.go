package models

import (
	"errors"
	"time"
)

var ErrProductNotFound = errors.New("Product not found")

// Product defines the structure for an API product
type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

// ProductsList is a collection of Product
type ProductsList []*Product

// GetProducts returns a list of products
func GetProducts() ProductsList {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func findProduct(id uint) (*Product, int, error) {
	for i, prod := range productList {
		if prod.ID == id {
			return prod, i, nil
		}
	}

	return nil, 0, ErrProductNotFound
}

func getNextID() uint {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func UpdateProduct(id uint, prod *Product) error {
	_, index, err := findProduct(id)
	if err != nil {
		return err
	}

	prod.ID = id
	productList[index] = prod
	return nil
}

// productList is a hard coded list of products for this
// example data source
var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffe",
		Price:       2.45,
		SKU:         "abc123",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and string coffe without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}
