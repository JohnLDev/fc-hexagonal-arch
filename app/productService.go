package app

import "fmt"

type ProductService struct {
	Persistance ProductPersistenceInterface
}

func (s ProductService) Get(id string) (ProductInterface, error) {
	return s.Persistance.Get(id)
}

func (s ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price

	if _, err := product.IsValid(); err != nil {
		return nil, err
	}

	return s.Persistance.Save(product)
}

func (s ProductService) Enable(product ProductInterface) (bool, error) {
	err := product.Enable()
	if err != nil {
		return false, err
	}
	fmt.Println(product)
	_, err = s.Persistance.Save(product)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s ProductService) Disable(product ProductInterface) (bool, error) {
	err := product.Disable()
	if err != nil {
		return false, err
	}

	_, err = s.Persistance.Save(product)
	if err != nil {
		return false, err
	}
	return true, nil
}
