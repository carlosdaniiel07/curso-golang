package main

import "fmt"

// O init é a 1a função a ser executada, antes mesmo da função main!
func init() {
	fmt.Println("Hello init")
}

func main() {
	fmt.Println("Hello main!")
}
