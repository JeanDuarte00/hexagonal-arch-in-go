package product

import (
	"errors"
)

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLE = "disabled"
	ENABLED = "enabled"
)

const (
	INVALID_PRICE           = "invalide price, must be greater than zero"
	CANT_DISABLE            = "price must be zero to desible"
	PRICE_GREATER_THAN_ZERO = "price must be greater than zero"
	INVALID_STATUS          = "status must be enabled or disabled"
)

type Product struct {
	ID     string
	Name   string
	Price  float64
	Status string
}

func (product *Product) IsValid() (bool, error) {
	if product.Status == "" {
		product.Status = DISABLE
	}
	if product.Status != ENABLED && product.Status != DISABLE {
		return false, errors.New(INVALID_STATUS)
	}
	if product.Price < 0 {
		return false, errors.New(PRICE_GREATER_THAN_ZERO)
	}
	return true, nil
}

func (product *Product) Enable() error {
	if product.Price <= 0 {
		return errors.New(INVALID_PRICE)
	}
	product.Status = ENABLED
	return nil
}

func (product *Product) Disable() error {
	if product.Price > 0 {
		return errors.New(CANT_DISABLE)
	}
	product.Status = DISABLE
	return nil
}

func (product *Product) GetId() string {
	return product.ID
}

func (product *Product) GetName() string {
	return product.Name
}

func (product *Product) GetPrice() float64 {
	return product.Price
}

func (product *Product) GetStatus() string {
	return product.Status
}
