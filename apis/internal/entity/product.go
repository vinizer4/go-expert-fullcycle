package entity

import (
	"errors"
	"github.com/vinizer4/go-expert-fullcycle/apis/pkg/entity"
	"time"
)

var (
	errIDIsRequired    = errors.New("id is required")
	errNameIsRequired  = errors.New("name is required")
	errPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("price must be greater than 0")
	ErrInvalidID       = errors.New("id must be a valid UUID")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return errIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}
	if p.Name == "" {
		return errNameIsRequired
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	if p.Price == 0 {
		return errPriceIsRequired
	}
	return nil
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}
