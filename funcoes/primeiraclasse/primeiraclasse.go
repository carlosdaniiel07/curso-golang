package main

import "fmt"

// Em Go as funções são "high order functions" (funções de primeira classe), portanto
// podem ser armazenadas em variáveis, asssim como no JavaScript, por exemplo
func main() {
	// Armazenando uma função na variável "soma'
	soma := func(a, b float64) float64 {
		return a + b
	}

	// Armazenando uma função na variável "sub'
	sub := func(a, b float64) float64 {
		return a - b
	}

	fmt.Println(soma(10, 20))
	fmt.Println(sub(10, 20))
}
