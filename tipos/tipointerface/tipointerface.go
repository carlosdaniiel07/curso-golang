package main

import "fmt"

type Curso struct {
	nome string
}

func main() {
	// Tipo genérico (any do TypeScript, object do C# e Java)
	var coisa interface {}

	coisa = 3

	fmt.Println(coisa)

	coisa = "Fulano"

	fmt.Println(coisa)

	coisa = false

	fmt.Println(coisa)
}
