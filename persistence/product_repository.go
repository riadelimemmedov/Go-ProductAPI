package persistence

import (
	"context"
	"errors"
	"fmt"
	"product-app/domain"
	"product-app/persistence/common"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
	AddProduct(product domain.Product) error
	GetProductById(productId int64) (domain.Product, error)
	DeleteProductById(productId int64) error
	UpdateProductPrice(productId int64, newPrice float32) error
}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{
		dbPool: dbPool,
	}
}

// !GetAllProducts
func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	ctx := context.Background()
	productRows, err := productRepository.dbPool.Query(ctx, "SELECT * FROM product")

	if err != nil {
		log.Error("Error while getting products", err)
	}
	return extractProductsFromRows(productRows)
}

// !GetAllProductsByStore
func (productRepository *ProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	ctx := context.Background()

	getProductsByStoreNameSql := `SELECT * FROM product WHERE store=$1`

	productRows, err := productRepository.dbPool.Query(ctx, getProductsByStoreNameSql, storeName)

	if err != nil {
		log.Error("Failed to execute query for getting products by store name")
	}
	return extractProductsFromRows(productRows)
}

// !AddProduct
func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	ctx := context.Background()

	insertProductSql := `INSERT INTO product (name,price,discount,store) VALUES ($1,$2,$3,$4)`

	addNewProduct, err := productRepository.dbPool.Exec(ctx, insertProductSql, product.Name, product.Price, product.Discount, product.Store)

	if err != nil {
		log.Error("Failed to add new product", err)
		return err
	} else {
		log.Info(fmt.Printf("Product add to database successfully: %s", addNewProduct))
	}
	return nil
}

// !GetProductById
func (productRepository *ProductRepository) GetProductById(productId int64) (domain.Product, error) {
	ctx := context.Background()

	getProductById := `SELECT * FROM product WHERE id=$1`

	queryRow := productRepository.dbPool.QueryRow(ctx, getProductById, productId)

	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	scanErr := queryRow.Scan(&id, &name, &price, &discount, &store)

	if scanErr != nil && scanErr.Error() == common.NOT_FOUND {
		return domain.Product{}, errors.New(fmt.Sprintf("Product not found with id %d", productId))
	}
	if scanErr != nil {
		return domain.Product{}, errors.New(fmt.Sprintf("Error while getting product with id %d", productId))
	}

	return domain.Product{
		Id:       id,
		Name:     name,
		Price:    price,
		Discount: discount,
		Store:    store,
	}, nil
}

// !DeleteProductById
func (productRepository *ProductRepository) DeleteProductById(productId int64) error {
	ctx := context.Background()

	_, getErr := productRepository.GetProductById(productId)

	if getErr != nil {
		return errors.New("Product not found")
	}

	deleteProductSql := "DELETE FROM product WHERE id=$1"

	_, err := productRepository.dbPool.Exec(ctx, deleteProductSql, productId)

	if err != nil {
		return errors.New(fmt.Sprintf("Error while delete product with id %d", productId))
	}
	log.Info("Product deleted successfully")
	return nil
}

// !UpdateProductPrice
func (productRepository *ProductRepository) UpdateProductPrice(productId int64, newPrice float32) error {
	ctx := context.Background()

	updateProductSql := `UPDATE product SET price=$1 WHERE id=$2`
	_, err := productRepository.dbPool.Exec(ctx, updateProductSql, newPrice, productId)

	if err != nil {
		return errors.New(fmt.Sprintf("Error while updating product with id %d", productId))
	}
	log.Info("Product price update successfully")
	return nil
}

// ?extractProductsFromRows
func extractProductsFromRows(productRows pgx.Rows) []domain.Product {
	var products = []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	for productRows.Next() {
		productRows.Scan(&id, &name, &price, &discount, &store)
		products = append(products, domain.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Discount: discount,
			Store:    store,
		})
	}
	return products
}
