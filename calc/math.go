package calc

import "errors"

//Add - Função para somar todos os valores passados por parâmetro e um código de erro
func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 2 {
		return sum, errors.New("Numero insuficiente de parametros, informar no minimo 2 números")
	} else {
		for _, num := range numbers {
			sum = sum + num
		}
	}

	return sum, nil
}
