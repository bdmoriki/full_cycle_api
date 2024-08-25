package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/bdmoriki/full_cycle_api/internal/entity"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	db, err := GetDBProduct()
	if err != nil {
		t.Error(err)
	}

	product, _ := entity.NewProduct("Teclado Logitech", 500)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	if err != nil {
		t.Error(err)
	}

	var productFound entity.Product
	err = db.First(&productFound).Where("id = ?", product.ID).Error
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, productFound)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
	assert.NotNil(t, product.CreatedAt, productFound.CreatedAt)
}

func TestFindAll(t *testing.T) {
	db, err := GetDBProduct()
	if err != nil {
		t.Error(err)
	}

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")

	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")

	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")

	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestFindById(t *testing.T) {
	db, err := GetDBProduct()
	if err != nil {
		t.Error(err)
	}

	product, _ := entity.NewProduct("Controle Xbox", 300)

	db.Create(product)

	productDB := NewProduct(db)
	productFound, err := productDB.FindById(product.ID.String())
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, productFound)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
	assert.NotNil(t, product.CreatedAt, productFound.CreatedAt)
}

func TestUpdateProduct(t *testing.T) {
	db, err := GetDBProduct()
	if err != nil {
		t.Error(err)
	}

	product, _ := entity.NewProduct("Xbox", 4000)

	db.Create(product)

	product.Name = "Playstation"
	product.Price = 3500

	productDB := NewProduct(db)
	err = productDB.Update(product)
	if err != nil {
		t.Error(err)
	}

	productFound, err := productDB.FindById(product.ID.String())
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, productFound)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := GetDBProduct()
	if err != nil {
		t.Error(err)
	}

	product, _ := entity.NewProduct("Nintendo Switch", 2000)

	db.Create(product)

	productDB := NewProduct(db)
	err = productDB.Delete(product.ID.String())
	if err != nil {
		t.Error(err)
	}

	productFound, _ := productDB.FindById(product.ID.String())

	assert.Equal(t, "", productFound.Name)
}
