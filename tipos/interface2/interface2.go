package main

import "fmt"

type Esportivo interface {
	ligarTurbo()
}

type Ferrari struct {
	modelo string
	turobLigado bool
	velocidadeAtual float64
}

func (f* Ferrari) ligarTurbo() {
	f.turobLigado = true
}

func main() {
	carro := Ferrari {
		modelo: "F40",
		turobLigado: false,
		velocidadeAtual: 200,
	}

	carro.ligarTurbo()

	var carro2 Esportivo = &Ferrari {
		modelo: "F40",
		turobLigado: false,
		velocidadeAtual: 200,
	}

	fmt.Println(carro)
	fmt.Println(carro2)
}
