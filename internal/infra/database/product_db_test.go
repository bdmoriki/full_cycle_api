package database

import (
	"testing"
	"time"

	"github.com/bdmoriki/full_cycle_api/internal/entity"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product := entity.Product{
		Name:  "Teclado Logitech",
		Price: 500,
		CreatedAt: time.Date(
			2009, 11, 17, 20, 34, 58, 651387237, time.UTC)}

	productDB := NewProduct(db)
	err = productDB.Create(&product)
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

func TestFindById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product := entity.Product{
		Name:  "Controle Xbox",
		Price: 300,
		CreatedAt: time.Date(
			2009, 9, 17, 20, 34, 58, 651387237, time.UTC)}

	productDB := NewProduct(db)
	err = productDB.Create(&product)
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
	assert.NotNil(t, product.CreatedAt, productFound.CreatedAt)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product := entity.Product{
		Name:  "Xbox",
		Price: 4000,
		CreatedAt: time.Date(
			2009, 12, 17, 20, 34, 58, 651387237, time.UTC)}

	productDB := NewProduct(db)
	err = productDB.Create(&product)
	if err != nil {
		t.Error(err)
	}

	product.Name = "Playstation"
	product.Price = 3500

	err = productDB.Update(&product)
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
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product := entity.Product{
		Name:  "Nintendo Switch",
		Price: 2000,
		CreatedAt: time.Date(
			2009, 13, 17, 20, 34, 58, 651387237, time.UTC)}

	productDB := NewProduct(db)
	err = productDB.Create(&product)
	if err != nil {
		t.Error(err)
	}

	err = productDB.Delete(product.ID.String())
	if err != nil {
		t.Error(err)
	}

	productFound, _ := productDB.FindById(product.ID.String())

	assert.Equal(t, "", productFound.Name)
}
