package dto

import "github.com/johnldev/go-hexagonal/app"

type Product struct {
	ID     string  `json:"id" valid:"uuidv4,optional"`
	Name   string  `json:"name" valid:"required"`
	Status string  `json:"status" valid:"required"`
	Price  float64 `json:"price" valid:"float,required"`
}

func (p *Product) Bind(product *app.Product) (*app.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}
	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status
	_, err := product.IsValid()
	if err != nil {
		return nil, err
	}

	return product, nil
}
