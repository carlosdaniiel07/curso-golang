package main

import "fmt"

type Exotico interface {
	ligarNitro()
}

type Esportivo interface {
	ligarTurbo()
}

// Composição de interfaces
type ExoticoEsportivo interface {
	Exotico
	Esportivo

	// Também é possíel adicionar mais métodos..
}

type BmwM3 struct {}
type Lamborghini struct {}

func (receiver BmwM3) ligarNitro()  {
	fmt.Println("Ativando nitro!")
}

func (receiver BmwM3) ligarTurbo()  {
	fmt.Println("Ativando turbo!")
}

func main() {
	var carro ExoticoEsportivo = BmwM3{}

	carro.ligarNitro()
	carro.ligarTurbo()
}


