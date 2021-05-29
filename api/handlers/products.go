package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Haizza1/api/models"
)

type Products struct {
	log *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lp := models.GetProducts()

	data, err := json.Marshal(&lp)
	if err != nil {
		http.Error(w, "Error while parsing response...", 500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
