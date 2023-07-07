package database_test

import (
	"database/sql"
	"testing"

	"github.com/johnldev/go-hexagonal/adapters/database"
	"github.com/johnldev/go-hexagonal/app"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestProductDb(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()
	suit := new(ProductDbTestSuit)
	suit.db = db

	suite.Run(t, suit)
}

type ProductDbTestSuit struct {
	suite.Suite
	db *sql.DB
}

func (s *ProductDbTestSuit) SetupTest() {
	s.db.Exec("create table if not exists products (id text, name text, status text, price real)")
}

func (s *ProductDbTestSuit) TestProductDb() {
	s.T().Run("Get", func(t *testing.T) {
		productDb := database.NewProductDb(s.db)
		procuct := app.NewProduct()
		procuct.Name = "Test"
		procuct.Price = 10

		productDb.Save(procuct)
		productService := app.ProductService{Persistance: productDb}
		productFound, err := productService.Get(procuct.ID)
		assert.Nil(s.T(), err)
		assert.Equal(s.T(), procuct.ID, productFound.GetID())
	})

	s.T().Run("Save", func(t *testing.T) {
		productDb := database.NewProductDb(s.db)
		productService := app.ProductService{Persistance: productDb}

		productSaved, err := productService.Create("Test", 10)
		assert.Nil(s.T(), err)
		assert.NotEmpty(s.T(), productSaved.GetID())

		productFound, err := productService.Get(productSaved.GetID())
		assert.Nil(s.T(), err)
		assert.Equal(s.T(), productSaved.GetID(), productFound.GetID())

		productTobeUpdated := app.Product{
			ID:     productFound.GetID(),
			Name:   "Test Updated",
			Price:  20,
			Status: productFound.GetStatus(),
		}

		productUpdated, err := productService.Persistance.Save(&productTobeUpdated)

		assert.Nil(s.T(), err)
		assert.Equal(s.T(), productUpdated.GetID(), productSaved.GetID())
		assert.Equal(s.T(), productUpdated.GetName(), "Test Updated")
		assert.Equal(s.T(), productUpdated.GetPrice(), 20.0)
	})
}
