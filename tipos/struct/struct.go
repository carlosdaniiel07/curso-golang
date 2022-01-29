package main

import "fmt"

// Struct: agrupamento de dados
type produto struct {
	nome string
	preco float64
	desconto float64
}

// Método: função com receiver (como se fosse um método da classe)
func (p produto) imprime() {
	fmt.Println("Nome:", p.nome)
	fmt.Println("Preço:", p.preco)
	fmt.Println("Desconto:", p.desconto)
}

// Método: função com receiver (como se fosse um método da classe)
func (p produto) valorLiquido() float64 {
	return p.preco - p.desconto
}

// Método: função com receiver (como se fosse um método da classe)
func (p produto) aplicaDesconto(percentual float64) float64 {
	return p.preco - (p.preco * percentual)
}

func main() {
	p := produto {
		nome:     "Ps5",
		preco:    1000,
		desconto: 100,
	}

	// Forma reduzida
	p2 := produto {"Ps4", 4000, 250}

	p.imprime()

	fmt.Println("Valor final: R$", p.valorLiquido())
	fmt.Println("Valor final com 10% de desconto:", p.aplicaDesconto(0.1))

	fmt.Println(p2)
}
