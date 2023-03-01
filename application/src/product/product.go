package product

import (
	"errors"
	"log"

	validator "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Product struct {
	id     string  `valid:"UUIDv4"`
	name   string  `valid:"required"`
	price  float64 `valid:"float,optional"`
	status string  `valid:"required"`
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

func Build(name string, price float64) Product {

	prod := Product{
		id:    uuid.NewV4().String(),
		name:  name,
		price: price,
	}
	err := prod.Enable()
	if err != nil {
		prod.Disable()
	}
	return prod
}

func init() {
	log.Println("Rodou init")
	validator.SetFieldsRequiredByDefault(true)
}

func (product *Product) IsValid() (bool, error) {

	if product.price < 0 {
		return false, errors.New(PRICE_GREATER_THAN_ZERO)
	}

	works, err := validator.ValidateStruct(product)

	if !works && err != nil {
		log.Println(INVALID_PRODUCT)
	}

	return works, err
}

func (product *Product) Enable() error {
	if product.price <= 0 {
		return errors.New(INVALID_PRICE)
	}
	product.status = ENABLED
	return nil
}

func (product *Product) Disable() error {
	if product.price > 0 {
		return errors.New(CANT_DISABLE)
	}
	product.status = DISABLE
	return nil
}

func (product *Product) GetId() string {
	return product.id
}

func (product *Product) GetName() string {
	return product.name
}

func (product *Product) GetPrice() float64 {
	return product.price
}

func (product *Product) GetStatus() string {
	return product.status
}
