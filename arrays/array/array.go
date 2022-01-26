package main

import "fmt"

func main() {
	// estrutura homogênea com tamanho fixo!
	var notas [4]float64

	notas[0], notas[1], notas[2], notas[3] = 10, 7, 7.5, 8.3

	fmt.Println(notas)
	fmt.Println(len(notas))
	fmt.Printf("Média do aluno: %.1f", calcularMedia(notas))
	fmt.Println()
	fmt.Printf("Média do aluno: %.1f", calcularMedia2(notas))
	fmt.Println()

	// inicializando array inline
	numeros := [...]int {20, 30, 40, 50, 60}
	fmt.Println(numeros)
}

func calcularMedia(notas [4]float64) float64 {
	tamanhoArray := len(notas)
	soma := 0.0

	for i := 0; i < tamanhoArray; i++ {
		nota := notas[i]
		soma += nota
	}

	return soma / float64(tamanhoArray)
}

// percorrendo array através do range
func calcularMedia2(notas [4]float64) float64 {
	soma := 0.0

	for index, nota := range notas {
		fmt.Println(index)
		soma += nota
	}

	return soma / float64(len(notas))
}