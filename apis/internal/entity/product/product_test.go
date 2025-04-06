package product

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 100)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 100, p.Price)
}

func TestProductWhenNameisRequired(t *testing.T) {
	p, err := NewProduct("", 100)
	assert.NotNil(t, err)
	assert.Equal(t, errNameIsRequired, err)
	assert.Nil(t, p)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Product 1", 0)
	assert.NotNil(t, err)
	assert.Equal(t, errPriceIsRequired, err)
	assert.Nil(t, p)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Product 1", -100)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
	assert.Nil(t, p)
}
