package main

import "fmt"

func imprimirX() func() {
	x := 10

	return func() {
		fmt.Println("Valor de x:", x)
	}
}

func main() {
	x := 20

	fmt.Println("Valor de x:", x)

	// A função lembra do contexto em que ela foi definida! (closure)
	imprimirX()()
}
