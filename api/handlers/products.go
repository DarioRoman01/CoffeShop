package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/Haizza1/api/models"
)

// Products is a http.Handler
type Products struct {
	log *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP entry point for the handler and staisfies the http.Handler interface
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		group := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(group) != 1 || len(group[0]) != 2 {
			http.Error(w, "Invalid URI", 400)
			return
		}

		id, err := strconv.Atoi(group[0][1])
		if err != nil {
			http.Error(w, "Unable to parse id as integer", 500)
			return
		}

		p.updateProduct(id, w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data store
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := models.GetProducts()

	err := lp.EncodeJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", 500)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	product := new(models.Product)
	if err := product.DecodeJSON(r.Body); err != nil {
		http.Error(w, "Unable to unmarshal json", 400) // bad request error
		return
	}

	models.AddProduct(product)
	if err := product.EncodeJSON(w); err != nil {
		http.Error(w, "Unable to marshal json", 500)
	}
}

func (p *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	product := new(models.Product)
	if err := product.DecodeJSON(r.Body); err != nil {
		http.Error(w, "Unable to unmarshal json", 400) // bad request error
		return
	}

	err := models.UpdateProduct(uint(id), product)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
