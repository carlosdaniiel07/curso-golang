package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println("Tipos básicos")

	fmt.Println(1, 2, 1000)

	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8, uint16, uint32, uint64
	var b byte = 255
	var c uint16 = 1223

	fmt.Println("O byte é", reflect.TypeOf(b))
	fmt.Println("O uint16 é", reflect.TypeOf(c))

	// com sinal... int8, int16, int32, int64
	d := math.MaxInt64

	fmt.Println("O valor máximo de um inteiro de 64 bits é", d)

	var e rune = 'a' // representa um mapeamento da tabela unicode (int32)

	fmt.Println("O rune é", reflect.TypeOf(e))
	fmt.Println("Valor da variável e:", e)

	// números reais (float32, float64)
	var x float32 = 49.99
	y := float32(49.99)

	fmt.Println("O tipo de X é", reflect.TypeOf(x))
	fmt.Println("O tipo literal de 49.99 é", reflect.TypeOf(49.99))
	fmt.Println(y)

	// boolean
	h := true

	fmt.Println("O tipo da variável h é", reflect.TypeOf(h))
	fmt.Print(h, !h)

	// string
	s1 := "Somente aspas duplas"
	s2 := `Dá para usar crase também!`
	s3 := `
		Múltiplas linhas,
		Sim é possível!
	`

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(len(s1))
	fmt.Println(len(s2))

	var teste bool = true
	var outroTeste string = "Outro teste"

	fmt.Println(teste)
	fmt.Println(outroTeste)
}
