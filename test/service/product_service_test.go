package service

import (
	"os"
	"product-app/service"
	"testing"
)

var productService service.IProductService

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}
