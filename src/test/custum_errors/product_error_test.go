package custum_errors

import (
	"github.com/kartikeya/product_catalog_DIY/src/main/custum_errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductError(t *testing.T) {
	productError := custum_errors.ProductError{
		Message: "error",
	}
	assert.Equal(t, "error", productError.Message)
}
