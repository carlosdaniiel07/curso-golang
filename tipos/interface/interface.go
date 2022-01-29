package main

import "fmt"

type Imprimivel interface {
	toString() string
}

// No Go não é necessário implementar a interface explicitamente, basta apenas
// ter um método com a mesma assinatura, no caso o método toString()
type Pessoa struct {
	nome string
	sobrenome string
	idade int
}

type Produto struct {
	nome string
	preco float64
}

func (p Pessoa) toString() string {
	return fmt.Sprintf("Nome: %s, sobrenome: %s, idade: %d", p.nome, p.sobrenome, p.idade)
}

func (p Produto) toString() string {
	return fmt.Sprintf("Nome: %s, preço: %.2f" , p.nome, p.preco)
}

func imprimir(objeto Imprimivel) {
	fmt.Println(objeto.toString())
}

func main() {
	pessoa := Pessoa {
		nome: "Fulano",
		sobrenome: "Ciclano",
		idade: 30,
	}
	produto := Produto {
		nome: "Produto",
		preco: 1250,
	}
	var pessoa2 Imprimivel = Pessoa {
		nome: "Fulano 2",
		sobrenome: "Ciclano 2",
		idade: 20,
	}

	imprimir(pessoa)
	imprimir(produto)
	imprimir(pessoa2)
}
