package service

import (
	"product-app/domain"
	"product-app/persistence"
)

type FakeProductRepository struct {
	products []domain.Product
}

// !NewFakeProductRepository
func NewFakeProductRepository(initialProducts []domain.Product) persistence.IProductRepository {
	return &FakeProductRepository{
		products: initialProducts,
	}
}

// !GetAllProducts
func (fakeRepository *FakeProductRepository) GetAllProducts() []domain.Product {
	return fakeRepository.products
}

// !GetAllProductsByStore
func (fakeRepository *FakeProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	return []domain.Product{}
}

// !AddProduct
func (fakeRepository *FakeProductRepository) AddProduct(product domain.Product) error {
	fakeRepository.products = append(fakeRepository.products, domain.Product{
		Id:       int64(len(fakeRepository.products)) + 1,
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	})
	return nil
}

// !GetProductById
func (fakeRepository *FakeProductRepository) GetProductById(productId int64) (domain.Product, error) {
	return domain.Product{}, nil
}

// !DeleteProductById
func (fakeRepository *FakeProductRepository) DeleteProductById(productId int64) error {
	return nil
}

// !UpdateProductPrice
func (fakeRepository *FakeProductRepository) UpdateProductPrice(productId int64, newPrice float32) error {
	return nil
}
