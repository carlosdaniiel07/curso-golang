package main

import (
	"fmt"
	m "math"
)

func main() {
	fmt.Println("Constantes e variáveis")

	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida (criação + atribuição)
	area := PI * m.Pow(raio, 2)

	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	println(a, b, c, d)

	var e, f bool = true, false
	g, h, i := 2, false, "epa!"

	println(e, f)
	println(g, h, i)
}
