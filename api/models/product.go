package models

import (
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator"
)

var (
	ErrProductNotFound = errors.New("Product not found")
	validate           = validator.New()
)

// Product defines the structure for an API product
type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

// Validate will use the validation tags to check
// if the product struct are valid we use strings
// builder to generate better error messages
func (p *Product) Validate() error {
	err := validate.Struct(p)
	validationErr := err.(validator.ValidationErrors)

	if len(validationErr) > 0 {
		var buff strings.Builder
		buff.WriteString("Missing fields: ")

		for i, err := range validationErr {
			buff.WriteString(err.Field())
			if i != len(validationErr)-1 {
				buff.WriteString(", ")
			}
		}
		return errors.New(buff.String())
	}

	return nil
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

// get the product with the given id
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

// update product data
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
