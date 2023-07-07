package cli

import (
	"encoding/json"
	"fmt"

	"github.com/johnldev/go-hexagonal/app"
)

func Run(service app.ProductServiceInterface, action, productId, productName string, price float64) (string, error) {
	var result string

	switch action {
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		_, err = service.Enable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s has been enabled", product.GetID())
		return result, nil

	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		_, err = service.Disable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s has been disabled", product.GetID())
		return result, nil

	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
		return result, nil

	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		jsonProduct, _ := json.MarshalIndent(product, "", "  ")
		result = string(jsonProduct)
		return result, nil
	}

}
