package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados:")

	for _, nome := range aprovados {
		fmt.Printf("- %s\n", nome)
	}
}

func main() {
	aprovados := []string { "Carlos", "Fulano", "Ciclano" }

	imprimirAprovados(aprovados...)
}
