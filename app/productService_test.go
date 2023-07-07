package app_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/johnldev/go-hexagonal/app"
	mock_app "github.com/johnldev/go-hexagonal/app/mocks"
	"github.com/stretchr/testify/assert"
)

func TestProductService(t *testing.T) {

	t.Run("Get", func(t *testing.T) {
		control := gomock.NewController(t)

		product := mock_app.NewMockProductInterface(control)
		persistence := mock_app.NewMockProductPersistenceInterface(control)
		persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

		service := app.ProductService{
			Persistance: persistence,
		}

		result, err := service.Get("abc")
		assert.Nil(t, err)
		assert.Equal(t, product, result)
	})

	t.Run("Create", func(t *testing.T) {
		control := gomock.NewController(t)

		product := mock_app.NewMockProductInterface(control)
		persistence := mock_app.NewMockProductPersistenceInterface(control)
		persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

		service := app.ProductService{
			Persistance: persistence,
		}

		result, err := service.Create("Test", 10)
		assert.Nil(t, err)
		assert.Equal(t, product, result)

		_, err = service.Create("Test", -10)
		assert.NotNil(t, err)
	})

	t.Run("Enable", func(t *testing.T) {
		control := gomock.NewController(t)

		product := mock_app.NewMockProductInterface(control)
		persistence := mock_app.NewMockProductPersistenceInterface(control)
		persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
		product.EXPECT().Enable().Return(nil).AnyTimes()

		service := app.ProductService{
			Persistance: persistence,
		}

		result, err := service.Enable(product)
		assert.Nil(t, err)
		assert.Equal(t, result, true)

	})

	t.Run("Disable", func(t *testing.T) {
		control := gomock.NewController(t)

		product := mock_app.NewMockProductInterface(control)
		persistence := mock_app.NewMockProductPersistenceInterface(control)
		persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
		product.EXPECT().Disable().Return(nil).AnyTimes()

		service := app.ProductService{
			Persistance: persistence,
		}

		result, err := service.Disable(product)
		assert.Nil(t, err)
		assert.Equal(t, result, true)

	})

}
