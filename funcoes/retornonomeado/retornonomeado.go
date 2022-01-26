package main

import "fmt"

// Retorno nomeado
func inverter(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1

	// Será retornado as variáveis "segundo" e "primeiro"
	return
}

func main() {
	p1, p2 := inverter(1, 2)

	fmt.Println(p1, p2)
}
