package product_test

import (
	"testing"
	product "training/go-hexagonal/application/src/product"

	"github.com/stretchr/testify/require"
)

func TestProduct_Disable(t *testing.T) {

	data := product.Build("xablau", 0)
	err := data.Disable()
	require.Nil(t, err)
	require.Equal(t, data.GetStatus(), product.DISABLE)

	data = product.Build("xablau", 20)
	err = data.Disable()
	require.NotNil(t, err)
	require.Equal(t, product.CANT_DISABLE, err.Error())
}

func TestProduct_Enable(t *testing.T) {

	data := product.Build("xablau", 0)
	err := data.Enable()
	require.NotNil(t, err)
	require.Equal(t, product.INVALID_PRICE, err.Error())

	data = product.Build("xablau", 20)
	err = data.Enable()
	require.Nil(t, err)
}

func TestProduct_IsValid(t *testing.T) {

	data := product.Build("xablau", 10)
	valid, err := data.IsValid()
	require.True(t, valid)
	require.Nil(t, err)

	data = product.Build("xablau", 0)
	valid, err = data.IsValid()
	require.False(t, valid)
	require.NotNil(t, err)
}
