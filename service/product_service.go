package service

import (
	"errors"
	"product-app/domain"
	"product-app/persistence"
	"product-app/service/model"
)

type IProductService interface {
	AllProducts() []domain.Product
	ProductsByStore(storeName string) []domain.Product
	Add(productCreate model.ProductCreate) error
	ProductById(productId int64) (domain.Product, error)
	DeleteById(productId int64) error
	UpdateProductPrice(productId int64, newPrice float32) error
}

type ProductService struct {
	productRepository persistence.IProductRepository
}

func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

// !Add
func (productService *ProductService) Add(productCreate model.ProductCreate) error {
	validateErr := validateProductCreate(productCreate)
	if validateErr != nil {
		return validateErr
	}
	return productService.productRepository.AddProduct(domain.Product{
		Name:     productCreate.Name,
		Price:    productCreate.Price,
		Discount: productCreate.Discount,
		Store:    productCreate.Store,
	})
}

// !DeleteById
func (productService *ProductService) DeleteById(productId int64) error {
	return productService.productRepository.DeleteProductById(productId)
}

// !ProductById
func (productService *ProductService) ProductById(productId int64) (domain.Product, error) {
	return productService.productRepository.GetProductById(productId)
}

// !UpdateProductPrice
func (productService *ProductService) UpdateProductPrice(productId int64, newPrice float32) error {
	return productService.productRepository.UpdateProductPrice(productId, newPrice)
}

// !AllProducts
func (productService *ProductService) AllProducts() []domain.Product {
	return productService.productRepository.GetAllProducts()
}

// !ProductsByStore
func (productService *ProductService) ProductsByStore(storeName string) []domain.Product {
	return productService.productRepository.GetAllProductsByStore(storeName)
}

// *validateProductCreate
func validateProductCreate(productCreate model.ProductCreate) error {
	if productCreate.Discount > 70.0 {
		return errors.New("Discount can not be greater than 70")
	}
	return nil
}
