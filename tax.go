package tax

import "errors"

// Criado uma interface de Repositório para emular o acesso ao banco de dados
type Repository interface {
	SaveTax(amount float64) error
}

// Emulando um save
func CalculateTaxAndSave(amount float64, repository Repository) error {
	// Calculando a taxa
	tax := CalculateTax2(amount)
	// Salvando no DB
	return repository.SaveTax(tax)
}

func CalculateTax2(amount float64) float64 {
	if amount <= 0 {
		return 0.0
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0
	}
	if amount >= 20000 {
		return 20.0
	}
	return 5.0
}

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0.0, errors.New("valor não pode ser menor ou igual a 0")
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0, nil
	}
	if amount >= 20000 {
		return 20.0, nil
	}
	return 5.0, nil
}
