package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(0)
	assert.Error(t, err, "valor não pode ser menor ou igual a 0")
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), "menor ou igual a 0")
}
