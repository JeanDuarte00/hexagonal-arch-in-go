package product_test

import (
	"testing"
	product "training/go-hexagonal/application/src/product"

	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {

	data := product.Product{}
	data.Name = "Xablau"
	data.Status = product.DISABLE

	data.Price = 10
	err := data.Enable()
	require.Nil(t, err)

	data.Price = 0
	err = data.Enable()
	require.NotNil(t, err)
	require.Equal(t, product.INVALID_PRICE, err.Error())
}

func TestProduct_Disable(t *testing.T) {

	data := product.Product{}
	data.Name = "Xablau"
	data.Status = product.ENABLED

	data.Price = 0
	err := data.Disable()
	require.Nil(t, err)
	require.Equal(t, data.Status, product.DISABLE)

	data.Price = 10
	err = data.Disable()
	require.NotNil(t, err)
	require.Equal(t, product.CANT_DISABLE, err.Error())

}
