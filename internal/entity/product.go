package entity

import (
	"errors"
	"time"

	"github.com/bdmoriki/full_cycle_api/pkg/entity"
)

var (
	errIDIsRequired    = errors.New("id is required")
	errInvalidID       = errors.New("id is invalid")
	errNameIsRequired  = errors.New("name is required")
	errPriceIsRequired = errors.New("price is required")
	errInvalidPrice    = errors.New("price is required")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()

	if err != nil {
		return product, err
	}

	return _, err
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return errIDIsRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return errInvalidID
	}

	if p.Name == "" {
		return errNameIsRequired
	}

	if p.Price == 0 {
		return errPriceIsRequired
	}

	if p.Price < 0 {
		return errInvalidPrice
	}

	return nil
}
