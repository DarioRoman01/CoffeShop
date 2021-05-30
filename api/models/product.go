package models

import (
	"encoding/json"
	"errors"
	"io"
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

func (p *Product) DecodeJSON(r io.Reader) error {
	p.CreatedAt = time.Now().UTC().String()
	p.UpdatedAt = time.Now().UTC().String()
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func (p *Product) EncodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(p)
}

// ProductsList is a collection of Product
type ProductsList []*Product

// EncodeJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *ProductsList) EncodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(p)
}

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
