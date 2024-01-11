package service

import (
	"os"
	"product-app/domain"
	"product-app/service"
	"product-app/service/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

var productService service.IProductService

func TestMain(m *testing.M) {
	initialProducts := []domain.Product{
		{
			Id:    1,
			Name:  "AirFryer",
			Price: 1000.0,
			Store: "ABC TECH",
		},
		{
			Id:    2,
			Name:  "Ütü",
			Price: 4000.0,
			Store: "ABC TECH",
		},
	}

	fakeProductReporitory := NewFakeProductRepository(initialProducts)
	productService = service.NewProductService(fakeProductReporitory)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func Test_ShouldGetAllProducts(t *testing.T) {
	t.Run("ShouldGetAllProducts", func(t *testing.T) {
		actualProducts := productService.AllProducts()
		assert.Equal(t, 2, len(actualProducts))
	})
}

func Test_WhenNoValidationErrorOccurred_ShouldAddProduct(t *testing.T) {
	t.Run("WhenNoValidationErrorOccurred_ShouldAddProduct", func(t *testing.T) {
		productService.Add(model.ProductCreate{
			Name:     "Ütü",
			Price:    2000.0,
			Discount: 50,
			Store:    "ABC TECH",
		})
		actualProducts := productService.AllProducts()
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, domain.Product{
			Id:       3,
			Name:     "Ütü",
			Price:    2000.0,
			Discount: 50,
			Store:    "ABC TECH",
		}, actualProducts[len(actualProducts)-1])
	})
}

// Discount can not be greater than 70
func Test_WhenDiscountIsHigherThan70_ShouldNotAddProduct(t *testing.T) {
	t.Run("WhenDiscountIsHigherThan70_ShouldNotAddProduct", func(t *testing.T) {
		err := productService.Add(model.ProductCreate{
			Name:     "Ütü",
			Price:    2000.0,
			Discount: 75,
			Store:    "ABC TECH",
		})
		actualProducts := productService.AllProducts()
		assert.Equal(t, 2, len(actualProducts))
		assert.Equal(t, "Discount can not be greater than 70", err.Error())
	})
}
