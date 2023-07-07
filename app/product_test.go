package app_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/johnldev/go-hexagonal/app"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {

	t.Run("Enable", func(t *testing.T) {
		assert := assert.New(t)
		product := app.Product{
			Name:   "Test",
			Price:  10,
			Status: app.DISABLED,
		}

		err := product.Enable()

		assert.Nil(err)

		product.Price = 0

		err = product.Enable()
		assert.Equal("the price must be greater than 0", err.Error())

	})

	t.Run("Disable", func(t *testing.T) {
		assert := assert.New(t)
		product := app.Product{
			Name:   "Test",
			Price:  0,
			Status: app.ENABLED,
		}

		err := product.Disable()

		assert.Nil(err)

		product.Price = 10

		err = product.Disable()
		assert.Equal("the price must be zero to disable the product", err.Error())

	})

	t.Run("IsValid", func(t *testing.T) {
		assert := assert.New(t)
		product := app.Product{
			Name:   "Test",
			Price:  10,
			Status: app.ENABLED,
			ID:     uuid.New().String(),
		}

		_, err := product.IsValid()

		assert.Nil(err)

		product.Status = "INVALID"

		_, err = product.IsValid()

		assert.Equal("the status must be enabled or disabled", err.Error())

		product.Status = app.DISABLED

		_, err = product.IsValid()

		assert.Nil(err)

		product.Price = -10

		_, err = product.IsValid()

		assert.Equal("the price must be greater or equal 0", err.Error())

	})

}
