package main

import "fmt"

func funcaoDefer() {
	fmt.Println("Essa linha é executada antes do return!")
}

func isAprovado(media float64) string {
	defer funcaoDefer()

	fmt.Println("Calculando situação do aluno..")

	if media >= 7 {
		return "Aprovado!"
	}

	return "Reprovado"
}

func main() {
	fmt.Println("Resultado:", isAprovado(10))
	fmt.Println("Resultado:", isAprovado(6))
	fmt.Println("Resultado:", isAprovado(7))
}
