package database

import (
	"github.com/bdmoriki/full_cycle_api/internal/entity"
	entityPKG "github.com/bdmoriki/full_cycle_api/pkg/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(p *entity.Product) error
	FindAll(page int, limit int, sort string) ([]entity.Product, error)
	FindById(id entityPKG.ID) (*entity.Product, error)
	Update(p *entity.Product) error
	Delete(id entityPKG.ID) error
}
