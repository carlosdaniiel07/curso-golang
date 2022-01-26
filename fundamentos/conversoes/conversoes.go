package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	z := x / float64(y)

	fmt.Println(z)

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)

	// cuidado... o método string() converte para a tabela ASCI/Unicode
	fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))
	fmt.Println("Teste " + fmt.Sprint(123))

	// string para int
	num, _ := strconv.Atoi("123")

	fmt.Println(num)

	// string para boolean
	b, _ := strconv.ParseBool("true")

	fmt.Println(b)
}
