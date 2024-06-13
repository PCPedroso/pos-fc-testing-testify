package tax

import (
	"errors"
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

func TestCalculateTaxAndSave(t *testing.T) {
	// Aqui é criado o repositorio onde vai emular o save no DB
	repository := &TaxRepositoryMock{}

	// Ao salvar taxa 10 não vai retornar erro na simulação
	repository.On("SaveTax", 10.0).Return(nil)

	// Ao tentar salvar taxa 0 vai retornar erro na simulação
	repository.On("SaveTax", 0.0).Return(errors.New("erro ao salvar taxa 0"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0, repository)
	assert.Error(t, err, "erro ao salvar taxa 0")

	repository.AssertExpectations(t)
}
