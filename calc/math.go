package calc

//Add - Função para somar todos os valores passados por parâmetro
func Add(numbers ...int) int {
	sum := 0

	for _, num := range numbers {
		sum = sum + num
	}

	return sum
}
