package main

import "fmt"

func fatorial(valor int) int {
	resultado := valor

	for i := valor - 1; i >= 1; i-- {
		resultado *= i
	}
	
	return resultado
}

func main() {
	resultado := fatorial(5)

	fmt.Println("Fatorial:", resultado)
}
