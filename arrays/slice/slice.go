package main

import (
	"fmt"
)

func main() {
	// O slice não é uma "cópia" do array, ele apenas mantém a
	// referência (ponteiro) para o array original!

	a := [3]int {1, 2, 3} // array
	b := []int {1, 2, 3} // slice

	// Pega o array do índice 0 ao índice 1
	c := a[0:2]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println()
	fmt.Println()

	notas := [4]float64 {10, 8, 9, 7}

	// Cria slice com tamanho inicial 10 e com capacidade 50. Quando a capacidade
	// "estourar" o programa irá automaticamente aumentar a capacidade do slice em runtime
	s := make([]float64, 10, 50)

	fmt.Println(notas)
	fmt.Println(s, len(s), cap(s))

	fmt.Println()
	fmt.Println()

	array1 := [3]int {1, 2, 3}
	var slice1 []int

	// Adicionando elementos ao slice através do método append
	slice1 = append(slice1, 10, 9, 8)
	slice2 := make([]int, 2)

	// Copiando os elementos do "slice1" para o "slice2". Será
	// copiado apenas 2 elementos visto que o tamanho do "slice2" é 2.
	// A função copy() não aumenta o tamanho do array assim como a função append()
	copy(slice2, slice1)

	fmt.Println(array1, slice1, slice2)
}