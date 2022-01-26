package main

import (
	"fmt"
)

func main() {
	i := 1

	// Criando variável do tipo ponteiro (asterisco)
	var ponteiro *int = nil

	// Pegando endereço de memória da variável 'i' e atribuindo a variável 'ponteiro'
	ponteiro = &i

	// Desferenciando (obtendo valor) o ponteiro e incrementando o seu valor
	*ponteiro++

	i++

	fmt.Println("Endereço de memória de 'i':", &i)
	fmt.Println("Endereço de memória de 'ponteiro':", ponteiro)
	fmt.Println("Valor de 'i':", i)
	fmt.Println("Valor de 'ponteiro':", *ponteiro)
}