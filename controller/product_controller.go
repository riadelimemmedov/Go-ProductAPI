package controller

import (
	"net/http"
	"product-app/controller/request"
	"product-app/controller/response"
	"product-app/service"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService *service.IProductService) *ProductController {
	return &ProductController{
		productService: *productService,
	}
}

func (productController *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/products/:id/", productController.ProductById)
	e.GET("/api/v1/products/", productController.AllProducts)
	e.POST("/api/v1/products/", productController.Add)
	e.PUT("/api/v1/products/:id/", productController.UpdateProductPrice)
	e.DELETE("/api/v1/products/:id/", productController.DeleteById)
}

func (productController *ProductController) ProductById(c echo.Context) error {
	return nil
}

func (productController *ProductController) AllProducts(c echo.Context) error {
	store := c.QueryParam("store")
	if len(store) == 0 {
		allProducts := productController.productService.AllProducts()
		return c.JSON(http.StatusOK, allProducts)
	}
	productWithGivenStore := productController.productService.ProductsByStore(store)
	return c.JSON(http.StatusOK, productWithGivenStore)
}

func (productController *ProductController) Add(c echo.Context) error {
	var addProductRequest request.AddProductRequest
	bindErr := c.Bind(&addProductRequest)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: bindErr.Error(),
		})
	}
	err := productController.productService.Add(addProductRequest.ToModel())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}
	return c.NoContent(http.StatusCreated)
}

func (product *ProductController) UpdateProductPrice(c echo.Context) error {
	return nil
}

func (product *ProductController) DeleteById(c echo.Context) error {
	return nil
}
