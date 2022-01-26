package main

import "fmt"

// Deve-se especificar a assinatura da função (parâmetros + tipo de retorno)
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	soma := func(p1, p2 int) int {
		return p1 + p2
	}

	sub := func(p1, p2 int) int {
		return p1 - p2
	}

	fmt.Println(exec(soma, 10, 20))
	fmt.Println(exec(sub, 90, 10))
}
