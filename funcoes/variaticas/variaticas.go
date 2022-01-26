package main

import "fmt"

// Função variática: parâmetros variáveis
func media(numeros ...float64) float64 {
	soma := 0.0

	for _, valor := range numeros {
		soma += valor
	}

	return soma / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(10, 10, 2, 1, 9, 7, 17))
}
