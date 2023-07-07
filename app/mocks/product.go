package mock_app

import (
	app "github.com/johnldev/go-hexagonal/app"
	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) Get(id string) (app.ProductInterface, error) {
	args := m.Called(id)
	return args.Get(0).(*app.Product), args.Error(1)
}

func (m *MockProductService) Create(name string, price float64) (app.ProductInterface, error) {
	args := m.Called(name, price)
	return args.Get(0).(*app.Product), args.Error(1)
}

func (m *MockProductService) Enable(product app.ProductInterface) (bool, error) {
	args := m.Called(product)
	return args.Bool(0), args.Error(1)

}

func (m *MockProductService) Disable(product app.ProductInterface) (bool, error) {
	args := m.Called(product)
	return args.Bool(0), args.Error(1)
}
