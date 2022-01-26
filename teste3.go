package main

import "fmt"

func increment(idade *int, valor int) {
	*idade += valor
}

func main() {
	idade := 22

	fmt.Println(idade)

	increment(&idade, 1)

	fmt.Println(idade)

	var ponteiro *int = nil

	ponteiro = &idade

	*ponteiro = 1

	fmt.Println(idade)
	fmt.Println(*ponteiro)
}
