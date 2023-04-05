package service

import (
	"fmt"
	"net/http"
	productRepository "sell-beauty-items/internal/products/repository"
)

type ProductService struct {
	Repository *productRepository.ProductRepository
}

func (p *ProductService) GetHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	alias := url.Path[len("/shop/"):]

	for v, k := range p.Repository.GetProducts() {
		if k.URL == alias {
			fmt.Println(v, k)
		}
	}
}

func (p *ProductService) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	p.Repository.GetProducts()
}
