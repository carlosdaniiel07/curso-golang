package main

import "fmt"

func main() {
	// O map é uma estrutura chave-valor, assim como o Dictionary no C#
	// Os mapas devem ser inicializados!
	aprovados := make(map[int]string)

	aprovados[50429693818] = "Carlos"
	aprovados[50429693819] = "Daniel"

	for chave, valor := range aprovados {
		fmt.Printf("CPF: %d - Nome: %s", chave, valor)
		fmt.Println()
	}

	fmt.Println(aprovados)

	funcionarios := map[string]float64 {
		"Fulano": 5000,
		"Ciclano": 7500,
	}

	for _, salario := range funcionarios{
		fmt.Printf("Salário: R$ %.2f", salario)
		fmt.Println()
	}
}
