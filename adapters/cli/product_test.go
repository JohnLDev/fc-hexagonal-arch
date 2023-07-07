package cli_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/johnldev/go-hexagonal/adapters/cli"
	"github.com/johnldev/go-hexagonal/app"
	mock_app "github.com/johnldev/go-hexagonal/app/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestSuit struct {
	suite.Suite
	product        *app.Product
	productService *mock_app.MockProductService
}

func (s *TestSuit) SetupTest() {
	product := app.NewProduct()
	product.Name = "Product 1"
	product.Price = 10.99

	s.product = product
	s.productService = new(mock_app.MockProductService)
}

func (s *TestSuit) TestCli_Create() {
	assert := assert.New(s.T())
	s.productService.On("Create", mock.Anything, mock.Anything).Return(s.product, nil).Once()
	result, err := cli.Run(s.productService, "create", "", "Product 1", 10.99)

	assert.Nil(err)
	assert.Equal(fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", s.product.GetID(), s.product.GetName(), s.product.GetPrice(), s.product.GetStatus()), result)
}

func (s *TestSuit) TestCli_Enable() {
	assert := assert.New(s.T())
	product := s.product
	s.productService.On("Get", product.ID).Return(product, nil).Once()
	s.productService.On("Enable", product).Return(true, nil).Once()
	result, err := cli.Run(s.productService, "enable", product.ID, "", 0)

	assert.Nil(err)
	assert.Equal(fmt.Sprintf("Product ID %s has been enabled", product.GetID()), result)
}

func (s *TestSuit) TestCli_Disable() {
	assert := assert.New(s.T())
	product := s.product
	s.productService.On("Get", product.ID).Return(product, nil).Once()
	s.productService.On("Disable", product).Return(true, nil).Once()
	result, err := cli.Run(s.productService, "disable", product.ID, "", 0)

	assert.Nil(err)
	assert.Equal(fmt.Sprintf("Product ID %s has been disabled", product.GetID()), result)
}

func (s *TestSuit) TestCli_Default_Get() {
	assert := assert.New(s.T())
	product := s.product
	s.productService.On("Get", product.ID).Return(product, nil).Once()
	result, err := cli.Run(s.productService, "", product.ID, "", 0)

	assert.Nil(err)
	jsonProduct, _ := json.MarshalIndent(product, "", "  ")
	assert.Equal(string(jsonProduct), result)
}

func TestProductCli(t *testing.T) {
	suite.Run(t, new(TestSuit))
}
