package repository

import "sell-beauty-items/internal/products/model"

var Products = []model.Product{
	{Name: "Product A", Price: 100, ImageURL: "/static/img/product-a.jpg"},
	{Name: "Product B", Price: 200, ImageURL: "/static/img/product-b.jpg"},
	{Name: "Product C", Price: 300, ImageURL: "/static/img/product-c.jpg"},
	{Name: "Product D", Price: 400, ImageURL: "/static/img/product-d.jpg"},
	{Name: "Product E", Price: 500, ImageURL: "/static/img/product-e.jpg"},
	{Name: "Product F", Price: 600, ImageURL: "/static/img/product-f.jpg"},
}

type ProductRepository struct{}

func (p *ProductRepository) GetProducts() []model.Product {
	return Products
}
