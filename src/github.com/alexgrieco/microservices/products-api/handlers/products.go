package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alexgrieco/microservices/products-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marchal json", http.StatusInternalServerError)
	}
	rw.Write(d)
}
