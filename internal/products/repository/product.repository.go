package repository

import "sell-beauty-items/internal/products/model"

var Products = []model.Product{
	{Name: "Product A", Price: 100, ImageURL: "/static/img/product-a.jpg", URL: "product-a"},
	{Name: "Product B", Price: 200, ImageURL: "/static/img/product-b.jpg", URL: "product-b"},
	{Name: "Product C", Price: 300, ImageURL: "/static/img/product-c.jpg", URL: "product-c"},
	{Name: "Product D", Price: 400, ImageURL: "/static/img/product-d.jpg", URL: "product-d"},
	{Name: "Product E", Price: 500, ImageURL: "/static/img/product-e.jpg", URL: "product-e"},
	{Name: "Product F", Price: 600, ImageURL: "/static/img/product-f.jpg", URL: "product-f"},
}

type ProductRepository struct{}

func (p *ProductRepository) GetProducts() []model.Product {
	return Products
}
