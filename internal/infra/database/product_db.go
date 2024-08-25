package database

import (
	"github.com/bdmoriki/full_cycle_api/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindAll(page int, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("createdAt " + sort).Find(products).Error
	} else {
		err = p.DB.Order("createdAt " + sort).Find(products).Error
	}

	return products, err
}

func (p *Product) FindById(id string) (*entity.Product, error) {
	product := entity.Product{}
	return &product, p.DB.First(&product).Where("id = ?", id).Error
}

func (p *Product) Update(product *entity.Product) error {
	return p.DB.Model(&entity.Product{}).Where("id = ?", product.ID).Updates(&entity.Product{Name: product.Name, Price: product.Price}).Error
}

func (p *Product) Delete(id string) error {
	return p.DB.Where("id = ?", id).Delete(&entity.Product{}).Error
}
