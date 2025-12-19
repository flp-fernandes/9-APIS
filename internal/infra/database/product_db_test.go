package database

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/flp-fernandes/9-APIS/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		t.Fatal(err)
	}

	err = db.Exec("DELETE FROM products").Error
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestCreateProduct(t *testing.T) {
	db := setupTestDB(t)

	product, err := entity.NewProduct("Product 1", 100.00)

	assert.NoError(t, err)

	productDB := NewProduct(db)

	err = productDB.Create(product)

	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db := setupTestDB(t)

	for i := range 23 {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i+1), rand.Float64()*100)

		assert.NoError(t, err)

		db.Create(product)
		assert.NoError(t, err)
	}

	productDB := NewProduct(db)

	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, products[0].Name, "Product 1")
	assert.Equal(t, products[9].Name, "Product 10")

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, products[0].Name, "Product 11")
	assert.Equal(t, products[9].Name, "Product 20")

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, products[0].Name, "Product 21")
	assert.Equal(t, products[2].Name, "Product 23")
}

func TestFindByIdProducts(t *testing.T) {
	db := setupTestDB(t)

	product, err := entity.NewProduct("Product 1", 100.00)
	assert.NoError(t, err)

	db.Create(product)
	productDB := NewProduct(db)

	productFound, err := productDB.FindById(product.ID.String())

	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	db := setupTestDB(t)

	product, err := entity.NewProduct("Product 1", 100.00)
	assert.NoError(t, err)

	db.Create(product)
	productDB := NewProduct(db)
	product.Name = "Product 2"

	err = productDB.Update(product)
	assert.NoError(t, err)

	product, err = productDB.FindById(product.ID.String())

	assert.NoError(t, err)
	assert.Equal(t, "Product 2", product.Name)
}

func TestDeleteProduct(t *testing.T) {
	db := setupTestDB(t)

	product, err := entity.NewProduct("Product 1", 100.00)
	assert.NoError(t, err)

	db.Create(product)
	productDB := NewProduct(db)
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	product, err = productDB.FindById(product.ID.String())
	assert.Error(t, err)
	assert.Nil(t, product)
}
